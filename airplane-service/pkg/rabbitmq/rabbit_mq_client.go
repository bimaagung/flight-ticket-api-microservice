package rabbitmq

import (
	"fmt"
	"log"
	"math"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQClient() (*ampq.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *ampq.Connection

	for {
		c, err := ampq.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off ...")
		time.Sleep(backoff)

		continue
	}

	return connection, nil


}