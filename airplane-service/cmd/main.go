package main

import (
	httphandler "airplane-service/internal/handler/http/v1"
	mysqlrepository "airplane-service/internal/repository/mysql_repository"
	"airplane-service/internal/usecase"
	"airplane-service/pkg/mysqldb"
	"airplane-service/pkg/rabbitmq"
	"log"
	"os"

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
	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetString(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASSWORD`)
	dbName := viper.GetString(`DB_NAME`)

	rabbitmqUser := viper.GetString(`RABBITMQ_USER`)
	rabbitmqPass := viper.GetString(`RABBITMQ_PASSWORD`)
	rabbitmqHost := viper.GetString(`RABBITMQ_HOST`)


	log.Println("Starting airplane service")

	// RabbitMQ connection
	rabbitConn, err := rabbitmq.NewRabbitMQClient(rabbitmqUser, rabbitmqPass, rabbitmqHost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	log.Println("Listening for and consuming RabbitMQ messages...")
	
	// connect to database
	conn := mysqldb.NewDBMysql(dbUser, dbPass, dbHost, dbPort, dbName)
	if conn == nil {
		log.Println(dbUser)
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

	port := viper.GetString("PORT")
	r.Run(port)

}