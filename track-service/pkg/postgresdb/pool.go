package postgresdb

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
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

func NewDBPostgres(dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode, dbTimezone, dbConnectTimeout string) *sql.DB{

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%s",
			dbHost,
			dbPort,
			dbUser,
			dbPass,
			dbName,
			dbSSLMode,
			dbTimezone,
			dbConnectTimeout,
		)

	for {
		connection, err := openDB(dns)
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