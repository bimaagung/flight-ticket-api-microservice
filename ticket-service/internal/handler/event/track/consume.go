package trackevent

import (
	"encoding/json"
	"fmt"
	"log"
	"ticket-service/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewTrackConsumer(conn *amqp.Connection)(Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	channel, err := conn.Channel()
	if err != nil {
		return Consumer{}, err
	}

	err = declareExchange(channel)

	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

type Consumer struct {
	conn *amqp.Connection
}

func (consumer *Consumer) Listen(topics string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		topics,
		"tracks_topic",
		false,
		nil,
	)

	if err != nil {
		return err
	}

	message, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		return err
	}

	
	for d := range message {
		// var payload domain.Track
		payload := &domain.Track{}
		err = json.Unmarshal(d.Body, &payload)
		if err != nil {
			fmt.Println("error: ", err)
		}

		log.Printf("arrival: %s, destination: %s", payload.Arrival, payload.Departure)
	}

    return nil
}
