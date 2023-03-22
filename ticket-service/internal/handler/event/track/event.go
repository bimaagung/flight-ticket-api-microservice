package trackevent

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *ampq.Channel) error {
	return ch.ExchangeDeclare(
		"tracks_topic",	// name
		"topic", 		// topic
		true,      		// durable?
        false,     		// auto-deleted?
		false,     		// internal?
        false,     		// noWait?
		nil, 			// arguments
	)
}

func declareRandomQueue(ch *ampq.Channel) (ampq.Queue, error) {
	return ch.QueueDeclare(
        "",     	// name?
        false,		// durable?
        false,		// delete when unused?
        false,		// exclusive?
        false,		// non-wait?
		nil,		// arguments?
	)
}