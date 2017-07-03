package pikaq

import "fmt"

func DeclareExchange(session *Session, exchange string, exchangeType string) (*Session, error) {
	if session == nil {
		return nil, fmt.Errorf("Invalid Session. Exchange can not be declared.")
	}

	err := session.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)

	if err != nil {
		return session, fmt.Errorf("DeclareExchange: %s", err)
	}

	return session, nil
}
