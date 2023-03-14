package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var counts int64

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBMysql() *sql.DB{

	dbHost := viper.Get(`DB_HOST`)
	dbPort := viper.Get(`DB_PORT`)
	dbUser := viper.Get(`DB_USER`)
	dbPass := viper.Get(`DB_PASSWORD`)
	dbName := viper.Get(`DB_NAME`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	val.Add("tls", "false")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Mysql not yet ready")
			counts++
		} else {
			log.Println("Connected to Mysql!")
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