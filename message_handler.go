package pikaq

import (
	"log"

	"github.com/streadway/amqp"
)

type Message amqp.Delivery
type Messages <-chan amqp.Delivery

type MessageHandler func(Messages, chan error)

var MessageLogger MessageHandler = func(msgs Messages, done chan error) {
	for m := range msgs {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(m.Body),
			m.DeliveryTag,
			m.Body,
		)
		m.Ack(false)
	}
	log.Printf("handle: msgs chan closed")
	done <- nil
}
