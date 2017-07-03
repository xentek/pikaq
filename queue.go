package pikaq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Queue amqp.Queue

func DeclareQueue(session *Session, queue string) (*Session, Queue, error) {
	var q amqp.Queue
	if session == nil {
		return session, Queue(q), fmt.Errorf("Invalid Session. Queue can not be declared.")
	}

	q, err := session.channel.QueueDeclare(
		queue, // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return session, Queue(q), fmt.Errorf("DeclareQueue: %s", err)
	}

	return session, Queue(q), err
}

func BindQueue(session *Session, queue Queue, key string, exchange string) (*Session, error) {
	var err error

	if session == nil {
		return session, fmt.Errorf("Invalid Session. Queue can not be bound to Exchange.")
	}

	err = session.channel.QueueBind(
		queue.Name, // name of the queue
		key,        // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("BindQueue: %s", err)
	}

	return session, nil
}
