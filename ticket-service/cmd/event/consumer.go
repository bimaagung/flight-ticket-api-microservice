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
	viper.AutomaticEnv()
	
	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
	}

	log.Println(viper.ConfigFileUsed())
}

func main() {
	dbHost 				:= viper.GetString(`DB_HOST`)
	dbPort 				:= viper.GetString(`DB_PORT`)
	dbUser 				:= viper.GetString(`DB_USER`)
	dbPass 				:= viper.GetString(`DB_PASSWORD`)
	dbName 				:= viper.GetString(`DB_NAME`)
	dbSSLMode 			:= viper.GetString(`DB_SSL_MODE`)
	dbTimezone 			:= viper.GetString(`DB_TIMEZONE`)
	dbConnectTimeout 	:= viper.GetString(`DB_CONNECT_TIMEOUT`)

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
	conn := postgresdb.NewDBPostgres(dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode, dbTimezone, dbConnectTimeout)
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