package rabbitmq

import (
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ExampleMain() {
	// Inisialisasi router Gin
	router := gin.Default()

	// Handler untuk endpoint /api/message
	router.POST("/api/message", func(c *gin.Context) {
		// Ambil data dari body request
		data := struct {
			Message string `json:"message"`
		}{}
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Kirim pesan ke RabbitMQ
		err = sendMessageToRabbitMQ(data.Message)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Berhasil mengirim pesan
		c.JSON(200, gin.H{"message": "success"})
	})

	// Handler untuk consumer RabbitMQ
	go consumeMessageFromRabbitMQ()

	// Jalankan server HTTP
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func sendMessageToRabbitMQ(message string) error {
	// Koneksi ke RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Deklarasi exchange
	err = ch.ExchangeDeclare(
		"my_exchange", // nama exchange
		"fanout",      // tipe exchange
		true,          // durable
		false,         // auto-delete
		false,         // internal
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	// Publish pesan ke exchange
	err = ch.Publish(
		"my_exchange", // nama exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func consumeMessageFromRabbitMQ() {
	// Koneksi ke RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
		"my_exchange", // nama exchange
		"fanout",      // tipe exchange
		true,          // durable
		false,         // auto-delete
		false,         // internal
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	// Deklarasi queue
	q, err := ch.QueueDeclare(
		"",    // nama queue (kosong = RabbitMQ akan generate nama acak)
		false, // durable
		false, // auto-delete
		true,  // exclusive (hanya bisa digunakan oleh consumer yang sama dengan channel yang sama)
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Bind queue ke exchange
	err = ch.QueueBind(
		q.Name,        // nama queue
		"",            // routing key (kosong karena exchange adalah fanout)
		"my_exchange", // nama exchange
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	// Mulai konsumsi pesan
	msgs, err := ch.Consume(
		q.Name, // nama queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Loop untuk membaca pesan dari channel
	for msg := range msgs {
		// Proses pesan yang diterima
		log.Printf("Received a message: %s", msg.Body)
	}
}
