package trackevent

import (
	"encoding/json"
	"ticket-service/domain"
	"ticket-service/internal/usecase"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
}

func NewTrackConsumer(conn *amqp.Connection)(*Consumer, error) {
	consumer := &Consumer{
		conn: conn,
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = declareExchange(channel)

	if err != nil {
		return nil, err
	}

	return consumer, nil
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
		"track_topic",
		false,
		nil,
	)

	if err != nil {
		return err
	}

	message, err := ch.Consume(topics, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range message {
			var payload domain.Track
			_ = json.Unmarshal(d.Body, &payload)

			go 
		}
	}()

	fmt.Printf("Waiting for messages [Exchange, Queue] [logs_topic, %s]", q.Name)
	<-forever

    return nil
}


