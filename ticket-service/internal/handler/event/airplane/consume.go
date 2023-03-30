package airplaneevent

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"ticket-service/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewAirplaneConsumer(conn *amqp.Connection, database *sql.DB, airplaneUseCase domain.AirplaneUseCase)(Consumer, error) {
	consumer := Consumer{
		conn: conn,
		DB: database,
		airplaneUseCase: airplaneUseCase,
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
	DB *sql.DB
	airplaneUseCase domain.AirplaneUseCase
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
		"airplanes_topic",
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for{
		message, err := ch.Consume(q.Name, "", true, false, false, false, nil)

		if err != nil {
			return err
		}

		for d := range message {
			var payload domain.Airplane

			err = json.Unmarshal(d.Body, &payload)
			
			if err != nil {
				fmt.Println("error: ", err)
			}

			fmt.Println(payload)

			id, err := consumer.airplaneUseCase.Add(&payload)

			if err != nil {
				log.Println("error: ", err)
			}
			
			log.Printf("success add airplane with id: %s", id)

		}
	}
}
