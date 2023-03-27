package main

import (
	"log"
	"os"
	httphandler "track-service/internal/handler/http/v1"
	postgresrepository "track-service/internal/repository/postgres_repository"
	"track-service/internal/usecase"
	"track-service/pkg/postgresdb"
	"track-service/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
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

	rabbitmqUser := viper.GetString(`RABBITMQ_USER`)
	rabbitmqPass := viper.GetString(`RABBITMQ_PASSWORD`)
	rabbitmqHost := viper.GetString(`RABBITMQ_HOST`)

	log.Println("Starting track service")

	// RabbitMQ connection
	rabbitConn, err := rabbitmq.NewRabbitMQClient(rabbitmqUser, rabbitmqPass, rabbitmqHost)
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	trackRepositoryPostgres := postgresrepository.NewTrackRepositoryPostgres(conn, rabbitConn)
	trackUseCase := usecase.NewTrackUseCase(trackRepositoryPostgres)
	trackHttpHandler := httphandler.NewTrackHandler(trackUseCase)

	trackHttpHandler.Route(r)

	port := viper.GetString("PORT")
	r.Run(port) 

}