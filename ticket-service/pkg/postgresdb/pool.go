package postgresdb

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

var counts int64

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBPostgres() *sql.DB{
	dbHost := viper.Get("DB_HOST")
	dbPort := viper.Get("DB_PORT")
	dbUser := viper.Get("DB_USER")
	dbPassword := viper.Get("DB_PASSWORD")
	dbName := viper.Get("DB_NAME")
	dbSSLMode := viper.Get("DB_SSLMODE")
	dbTimezone := viper.Get("DB_TIMEZONE")
	dbConnectTimeout := viper.Get("DB_CONNECT_TIMEOUT")

	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%s", 
						dbHost, 
						dbPort,
						dbUser,
						dbPassword,
						dbName,
						dbSSLMode,
						dbTimezone,
						dbConnectTimeout) 

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}