package pikaq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Session struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewSession(url string) (*Session, error) {
	var (
		c   *amqp.Connection
		ch  *amqp.Channel
		err error
	)

	c, err = amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("Closing: %s", <-c.NotifyClose(make(chan *amqp.Error)))
	}()

	ch, err = c.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	return &Session{conn: c, channel: ch}, nil
}
