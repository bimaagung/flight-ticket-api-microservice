package main

import (
	"fmt"
	"log"
	httphandler "ticket-service/internal/handler/http/v1"
	postgresrepository "ticket-service/internal/repository/postgres_repository"
	"ticket-service/internal/usecase"
	"ticket-service/pkg/postgresdb"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	log.Println("Starting ticket service")

	// connect to database
	conn := postgresdb.NewDBPostgres()
	if conn == nil {
		log.Panic("Can't connect to database")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ticketRepositoryPostgres := postgresrepository.NewTicketPostgresRepository(conn)
	ticketUseCase := usecase.NewTicketUseCase(ticketRepositoryPostgres)
	ticketHttpHandler := httphandler.NewTicketHandler(ticketUseCase)

	ticketHttpHandler.Route(r)

	port := fmt.Sprintf(":%s", viper.Get("PORT"))
	r.Run(port) 

}