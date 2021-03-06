# pikaq [![Build Status](https://travis-ci.org/xentek/pikaq.svg?branch=master)](https://travis-ci.org/xentek/pikaq) [![GoDoc](https://godoc.org/github.com/xentek/pikaq?status.svg)](https://godoc.org/github.com/xentek/pikaq) [![Go Report Card](https://goreportcard.com/badge/github.com/xentek/pikaq)](https://goreportcard.com/report/github.com/xentek/pikaq) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xentek/pikaq/blob/master/LICENSE) 

![PikaQ, I choose you!](http://xentek-images.s3.amazonaws.com/pikachu-and-ash.png "PikaQ, I choose you!")

#### __PikaQ, I choose YOU!__

PikaQ makes working with Rabbit MQ _even_ easier. It provides a thin layer around [streadway/amqp](https://github.com/streadway/amqp), the defacto standard golang library for publishing and consuming Rabbit MQ messages. This package makes writing your own high performance message-based services fun and type safe.

> The name "pika" is used for any member of the Ochotonidae, a family within the order of lagomorphs, which also includes the Leporidae (rabbits and hares).
> __-- [Wikipedia](https://en.wikipedia.org/wiki/Pika)__

---

## Install

    go get https://github.com/xentek/pikaq

## Import

    import "github.com/xentek/pikaq"

## Usage

The simplest consumer you can create just logs and _acks_ each message it recieves. The example below uses the built in `LoggingHandler`. 

For your consumer, write a function with the following signature: `func(pikaq.Messages, chan error)`, and pass it as the last argument of `pikaq.LoggingHandler`. Refer to the [examples](_examples) for more information.

```golang
package main

import (
	"log"

	"github.com/xentek/pikaq"
)

func main() {
	c, err := pikaq.NewConsumer("amqp://localhost:5672", "amq.direct", "direct", "example-queue", "routing-key", "example", pikaq.LoggingHandler)
	if err != nil {
		log.Fatalf("New Consumer Error: %s", err)
	}
	log.Printf("Started Consumer: %s", c.tag.Tag())
}
```
## Contributing

Refer to our [Contributor's Guide](CONTRIBUTING.md) to learn how you can participate in this project.

## More Info

  - [GoDoc](https://godoc.org/github.com/xentek/pikaq)
  - [Wiki](https://github.com/xentek/pikaq/wiki)
  - [Examples](_examples)
