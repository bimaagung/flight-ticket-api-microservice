package main

import (
	"fmt"
	"log"
	httphandler "track-service/internal/handler/http/v1"
	postgresrepository "track-service/internal/repository/postgres_repository"
	"track-service/internal/usecase"
	"track-service/pkg/postgresdb"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	log.Println("Starting track service")

	dbPostgres := postgresdb.NewDBPostgres()

	trackRepositoryPostgres := postgresrepository.NewTrackRepositoryPostgres(dbPostgres)
	trackUseCase := usecase.NewTrackUseCase(trackRepositoryPostgres)
	trackHttpHandler := httphandler.NewTrackHandler(trackUseCase)

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

	trackHttpHandler.Route(r)

	port := fmt.Sprintf(":%s", viper.Get("PORT"))
	r.Run(port) 
}