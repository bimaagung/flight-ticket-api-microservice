package trackevent

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"ticket-service/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewTrackConsumer(conn *amqp.Connection, database *sql.DB, trackUseCase domain.TrackUseCase)(Consumer, error) {
	consumer := Consumer{
		conn: conn,
		DB: database,
		trackUseCase: trackUseCase,
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
	trackUseCase domain.TrackUseCase
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

	for{
		message, err := ch.Consume(q.Name, "", true, false, false, false, nil)

		if err != nil {
			return err
		}

		for d := range message {
			var payload domain.Track

			err = json.Unmarshal(d.Body, &payload)
			
			if err != nil {
				fmt.Println("error: ", err)
			}

			fmt.Println(payload)

			id, err := consumer.trackUseCase.Add(&payload)
			// var ID uuid.UUID = uuid.New()
			// query := `insert into tracks (id, arrival, departure, long_flight) values ($1, $2, $3, $4) returning id`

			// err := consumer.DB.QueryRow(query,
			// 	ID,
			// 	payload.Arrival,
			// 	payload.Departure,
			// 	payload.LongFlight,
			// ).Scan(&ID)

			if err != nil {
				log.Println("error: ", err)
			}
			
			log.Printf("success add track with id: %s", id)
		}
	}
}
