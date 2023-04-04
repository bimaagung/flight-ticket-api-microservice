package es

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var counts int64

func NewESClient(esHost, esPort string) *elasticsearch.Client {
	
	dsn := fmt.Sprintf("http://%s:%s", esHost, esPort)

	cfg := elasticsearch.Config{
		Addresses: []string{
			dsn,
		},
	}
	
	for {
		es, err := elasticsearch.NewClient(cfg)

		if err != nil {
			log.Println("Elasticsearch not yet ready")
		}else{
			log.Println("Connected to Elasticsearch!")
			return es
		}

		if counts > 10 {
			log.Println(err)
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}