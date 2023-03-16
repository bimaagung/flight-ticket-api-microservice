package main

import (
	httphandler "airplane-service/internal/handler/http/v1"
	mysqlrepository "airplane-service/internal/repository/mysql_repository"
	"airplane-service/internal/usecase"
	"airplane-service/pkg/mysqldb"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	log.Println("Starting airplane service")

	dbMysql := mysqldb.NewDBMysql()

	airplaneRepositoryPostgres := mysqlrepository.NewAirplaneRepositoryMysql(dbMysql)
	airplaneUseCase := usecase.NewAirplaneUseCase(airplaneRepositoryPostgres)
	airplaneHttpHandler := httphandler.NewAirplaneHandler(airplaneUseCase)

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

	airplaneHttpHandler.Route(r)

	port := fmt.Sprintf(":%s", viper.Get("PORT"))
	r.Run(port) 

}