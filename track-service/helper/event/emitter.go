package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	// setup
	channel, err := emitter.connection.Channel()
	
	if err != nil {
		return Emitter{},err
	}

	defer channel.Close()

	err = declareExchange(channel)
	
	if err != nil {
		return Emitter{}, err
	}	

	return emitter, nil
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
            Body:        []byte(event),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

