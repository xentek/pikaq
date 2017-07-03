package main

import (
	"log"

	"github.com/macandmia/pikaq"
)

func ExampleMessageHandler(msgs Messages, done chan error) {
	// range over the messages
	for m := range msgs {
		// do something with each message recieved
		// in this case, we're just logging some info about the message...
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(m.Body),
			m.DeliveryTag,
			m.Body,
		)
		// Ack the message
		m.Ack(false)
	}
	// send status on the "done" chan
	// if there is an issue, send the error to the "done" chan
	done <- nil
}

func main() {
	c, err := pikaq.NewConsumer("amqp://localhost:5672", "amq.direct", "direct", "example-queue", "routing-key", "example", ExampleMessageHandler)
	if err != nil {
		log.Fatalf("New Consumer Error: %s", err)
	}
	log.Printf("Started Consumer: %s", c.tag.Tag())
}
