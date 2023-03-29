package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessageFromRabbitMQ() {
	// Koneksi ke RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ server: %v", err)
	}
	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Deklarasi exchange
	err = ch.ExchangeDeclare(
		"tracks_topic",	// name
		"topic", 		// topic
		true,      		// durable?
        false,     		// auto-deleted?
		false,     		// internal?
        false,     		// noWait?
		nil, 			// arguments
	)


	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	// Deklarasi queue
	q, err := ch.QueueDeclare(
		"",     	// name?
        false,		// durable?
        false,		// delete when unused?
        false,		// exclusive?
        false,		// non-wait?
		nil,		// arguments?
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Bind queue ke exchange
	err = ch.QueueBind(
		q.Name,        // nama queue
		"track.INFO",            // routing key (kosong karena exchange adalah fanout)
		"tracks_topic", // nama exchange
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	// Mulai konsumsi pesan
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Loop untuk membaca pesan dari channel
	for msg := range msgs {
		// Proses pesan yang diterima
		log.Printf("Received a message: %s", msg.Body)
	}
}