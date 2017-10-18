package pikaq

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	session *Session
	tag     *ConsumerTag
	queue   Queue
	done    chan error
}

func NewConsumer(url string, exchange string, exchangeType string, queue string, key string, name string, handler MessageHandler) (*Consumer, error) {
	var err error
	tag := NewConsumerTag(name)

	log.WithFields(log.Fields{"consumer": tag.Tag(), "url": url}).Debug("Establishing Connection To Server...")
	s, err := NewSession(url)
	if err != nil {
		log.WithFields(log.Fields{"consumer": tag.Tag()}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"consumer": tag.Tag()}).Info("Session Created.")

	c := &Consumer{session: s, tag: tag}
	c.session, err = DeclareExchange(c.session, exchange, exchangeType)
	if err != nil {
		log.WithFields(log.Fields{"consumer": tag.Tag(), "exchange": exchange, "type": exchangeType}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"consumer": tag.Tag(), "exchange": exchange, "exchange_type": exchangeType}).Info("Exchange Declared.")

	c.session, c.queue, err = DeclareQueue(c.session, queue)
	if err != nil {
		log.WithFields(log.Fields{"consumer": tag.Tag(), "queue": queue}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"consumer": tag.Tag(), "queue": queue}).Info("Queue Declared.")

	c.session, err = BindQueue(c.session, c.queue, key, exchange)
	if err != nil {
		log.WithFields(log.Fields{"consumer": tag.Tag(), "queue": queue, "total_messages": c.queue.Messages, "total_consumers": c.queue.Consumers, "exchange": exchange, "exchange_type": exchangeType, "routing": key}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"consumer": tag.Tag(), "queue": queue, "total_messages": c.queue.Messages, "total_consumers": c.queue.Consumers, "exchange": exchange, "exchange_type": exchangeType, "routing_key": key}).Info("Queue Bound To Exchange.")

	msgs, err := c.Start()
	if err != nil {
		log.WithFields(log.Fields{"consumer": tag.Tag()}).Error(err)
		return nil, err
	}

	go handler(msgs, c.done)
	return c, err
}

func (c *Consumer) Start() (Messages, error) {
	var (
		msgs Messages
		err  error
	)

	c.session.channel.Qos(
		2,    // number of messages to send
		0,    // number of bytes to send (0 = no limit)
		true, // settings apply to all consumers on this channel
	)

	msgs, err = c.session.channel.Consume(
		c.queue.Name, // name
		c.tag.Tag(),  // consumerTag,
		false,        // noAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("consumer.Start: %s", err)
	}

	return msgs, nil
}

func (c *Consumer) Stop() error {
	var err error
	// will close() the deliveries channel
	err = c.session.channel.Cancel(c.tag.Tag(), true)
	if err != nil {
		return fmt.Errorf("consumer.Stop - Cancel Failed: %s", err)
	}

	err = c.session.conn.Close()
	if err != nil {
		return fmt.Errorf("consumer.Stop - connection Close Error: %s", err)
	}

	defer log.Printf("Consumer Stopped: %q", c.tag.Tag())

	// wait for handle() to exit
	return <-c.done
}

func (c *Consumer) Info() string {
	return c.tag.Tag()
}
