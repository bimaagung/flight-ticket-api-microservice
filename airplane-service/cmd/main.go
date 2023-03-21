package main

import (
	httphandler "airplane-service/internal/handler/http/v1"
	mysqlrepository "airplane-service/internal/repository/mysql_repository"
	"airplane-service/internal/usecase"
	"airplane-service/pkg/mysqldb"
	"airplane-service/pkg/rabbitmq"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	log.Println("Starting airplane service")

	// RabbitMQ connection
	rabbitConn, err := rabbitmq.NewRabbitMQClient()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	log.Println("Listening for and consuming RabbitMQ messages...")
	
	// connect to database
	conn := mysqldb.NewDBMysql()
	if conn == nil {
		log.Panic("Can't connect to database")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	airplaneRepositoryPostgres := mysqlrepository.NewAirplaneRepositoryMysql(conn, rabbitConn)
	airplaneUseCase := usecase.NewAirplaneUseCase(airplaneRepositoryPostgres)
	airplaneHttpHandler := httphandler.NewAirplaneHandler(airplaneUseCase)


	airplaneHttpHandler.Route(r)

	port := fmt.Sprintf(":%s", viper.Get("PORT"))
	r.Run(port) 

}