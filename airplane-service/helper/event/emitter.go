package event

import (
	"airplane-service/domain"
	"encoding/json"
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

func (e *Emitter) PushToQueue(payload *domain.Airplane, severity string) error {
	channel, err := e.connection.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	log.Println("Pushing to channel")

	log.Println(payload)

	j , err := json.MarshalIndent(&payload, "", "\t")

	if err != nil {
		return err
	}
	
	err = channel.Publish(
		"airplanes_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
            Body:        []byte(string(j)),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

