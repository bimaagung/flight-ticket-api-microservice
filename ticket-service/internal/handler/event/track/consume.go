package trackevent

import (
	"encoding/json"
	"fmt"
	"log"
	"ticket-service/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewTrackConsumer(conn *amqp.Connection, trackUseCase domain.TrackUseCase)(Consumer, error) {
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
	TrackUseCase domain.TrackUseCase
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

	forever := make(chan bool)
	go func() {
		for d := range message {
			// var payload domain.Track
			payload := &domain.Track{}
			err = json.Unmarshal(d.Body, &payload)
			if err != nil {
				fmt.Println("error: ", err)
			}

			fmt.Printf("arrival: %s, destination: %s", payload.Arrival, payload.Departure)

			if err := d.Ack(true); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
			
			//go consumer.TrackUseCase.Add(payload)
		}
	}()

	fmt.Printf("Waiting for messages [Exchange, Queue] [tracks_topic, %s]", q.Name)
	<-forever

    return nil
}
