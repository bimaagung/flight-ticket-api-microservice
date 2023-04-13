package es

import (
	"fmt"
	"log"
	"strings"
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

func AddIndex(es *elasticsearch.Client, index string) (string, error) {
    mapping :=
	`{
		"settings": {
			"index": {
				"number_of_shards": 5,
				"number_of_replicas": 2 
			},
			"analysis": {
				"analyzer": {
					"my_analyzer": {
						"tokenizer": "my_tokenizer"
					}
				},
				"tokenizer": {
					"my_tokenizer": {
						"type": "ngram",
						"min_gram": 3,
						"max_gram": 3
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"track.arrival": {
					"type": "text",
					"analyzer": "my_analyzer"
				},
				"track.departure": {
					"type": "text",
					"analyzer": "my_analyzer"
				}
			}
		}
	}`


	res, err := es.Indices.Create(
		index,
		es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)

	stmt := fmt.Sprintf("index %s is already", index)

	if res.StatusCode == 400 {
		return stmt, nil
	}

	if err != nil {
		return "", err
	}

	stmt = fmt.Sprintf("success created index %s ", index)
	return stmt, nil
}
