package main

import (
	"log"
	"os"
	trackevent "ticket-service/internal/handler/event/track"
	postgresrepository "ticket-service/internal/repository/postgres_repository"
	"ticket-service/internal/usecase"
	"ticket-service/pkg/postgresdb"
	"ticket-service/pkg/rabbitmq"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	log.Println("Starting ticket service")

	// RabbitMQ connection
	rabbitConn, err := rabbitmq.NewRabbitMQClient()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	log.Println("Listening for and consuming RabbitMQ messages...")


	// connect to database
	conn := postgresdb.NewDBPostgres()
	if conn == nil {
		log.Panic("Can't connect to database")
	}

	// Track
	trackRepositoryPostgres := postgresrepository.NewTrackRepositoryPostgres(conn)
	trackUseCase := usecase.NewTrackUseCase(trackRepositoryPostgres)

	// create consumer
	consumer, err := trackevent.NewTrackConsumer(rabbitConn, trackUseCase)

	if err != nil {
		log.Panic("Can't create consumer")
	}

	err = consumer.Listen("track.INFO")

	if err != nil {
		log.Println(err)
	}


}