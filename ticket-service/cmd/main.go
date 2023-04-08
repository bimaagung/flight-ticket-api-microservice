package main

import (
	"log"
	"os"
	airplaneevent "ticket-service/internal/handler/event/airplane"
	trackevent "ticket-service/internal/handler/event/track"
	httphandler "ticket-service/internal/handler/http/v1"
	esrepository "ticket-service/internal/repository/es_repository"
	postgresrepository "ticket-service/internal/repository/postgres_repository"
	"ticket-service/internal/usecase"
	"ticket-service/pkg/es"
	"ticket-service/pkg/postgresdb"
	"ticket-service/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
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

	rabbitmqUser := viper.GetString(`RABBITMQ_USER`)
	rabbitmqPass := viper.GetString(`RABBITMQ_PASSWORD`)
	rabbitmqHost := viper.GetString(`RABBITMQ_HOST`)

	esHost := viper.GetString(`ELASTICSEARCH_HOST`)
	esPort := viper.GetString(`ELASTICSEARCH_PORT`)

	log.Println("Starting ticket service")


	// connect to database
	conn := postgresdb.NewDBPostgres(dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode, dbTimezone, dbConnectTimeout)
	if conn == nil {
		log.Panic("Can't connect to database")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// connect to elasticsearch
	esClient := es.NewESClient(esHost, esPort)

	msg, err := es.AddIndex(esClient, "ticket")

	if err != nil {
		log.Println(err)
	}else{
		log.Println(msg)
	}

	// Track
	trackRepositoryPostgres := postgresrepository.NewTrackRepositoryPostgres(conn)
	trackUseCase := usecase.NewTrackUseCase(trackRepositoryPostgres)
	
	// Airplane
	airplaneRepositoryPostgres := postgresrepository.NewAirplaneRepositoryPostgres(conn)
	airplaneUseCase := usecase.NewAirplaneUseCase(airplaneRepositoryPostgres)

	// Ticket
	ticketRepositoryES := esrepository.NewTicketESRepository(esClient)
	ticketRepositoryPostgres := postgresrepository.NewTicketPostgresRepository(conn)
	ticketUseCase := usecase.NewTicketUseCase(ticketRepositoryPostgres, trackRepositoryPostgres, airplaneRepositoryPostgres, ticketRepositoryES)
	ticketHttpHandler := httphandler.NewTicketHandler(ticketUseCase)
	
	ticketHttpHandler.Route(r)


	// RabbitMQ connection
	rabbitConn, err := rabbitmq.NewRabbitMQClient(rabbitmqUser, rabbitmqPass, rabbitmqHost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	log.Println("Listening for and consuming RabbitMQ messages...")

	
	go func() {
		consumer, err := trackevent.NewTrackConsumer(rabbitConn, conn, trackUseCase)
		if err != nil {
			log.Println(err)
		}

		err = consumer.Listen("track.INFO")
		
		if err != nil {
			log.Println(err)
		}
	}()
	
	go func() {
		consumer, err := airplaneevent.NewAirplaneConsumer(rabbitConn, conn, airplaneUseCase)
		if err != nil {
			log.Println(err)
		}

		err = consumer.Listen("airplane.INFO")
		
		if err != nil {
			log.Println(err)
		}
	}()

	port := viper.GetString("PORT")
	r.Run(port) 

}