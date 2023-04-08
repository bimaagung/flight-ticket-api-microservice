package esrepository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"ticket-service/domain"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func NewTicketESRepository(esClient *elasticsearch.Client) domain.TicketRepositoryElasticsearch {
	return &ticketRepositoryES{
		ESClient: esClient,
		DBTimeout: time.Second * 3,
	}
}

type ticketRepositoryES struct {
	ESClient *elasticsearch.Client
	DBTimeout time.Duration
}

func (repository *ticketRepositoryES) Insert(idTicket string, ticket *domain.TicketRes) error {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	data, err := json.Marshal(ticket)

	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index: "ticket",
		DocumentID: idTicket,
		Body: bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(ctx, repository.ESClient)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 400 {
		return errors.New(res.String())
	}

	log.Printf("status_code: %d, success created ticket in elasticsearch",res.StatusCode)

	return nil
}

func (repository *ticketRepositoryES) Update(idTicket string, ticket *domain.TicketRes) error {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	data, err := json.Marshal(ticket)

	if err != nil {
		return err
	}

	req := esapi.UpdateRequest{
		Index: "ticket",
		DocumentID: idTicket,
		Body: bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, data))),
		Refresh: "true",
	}

	res, err := req.Do(ctx, repository.ESClient)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 400 {
		return errors.New(res.String())
	}

	log.Printf("status_code: %d, success updated ticket in elasticsearch",res.StatusCode)

	return nil
}

func (repository *ticketRepositoryES) Delete(idTicket string) error {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	req := esapi.DeleteRequest{
		Index: "ticket",
		DocumentID: idTicket,
		Refresh: "true",
	}

	res, err := req.Do(ctx, repository.ESClient)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 400 {
		return errors.New(res.String())
	}

	log.Printf("status_code: %d, success deleted ticket in elasticsearch",res.StatusCode)

	return nil
}

