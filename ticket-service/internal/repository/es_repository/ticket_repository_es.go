package esrepository

import (
	"bytes"
	"context"
	"encoding/json"
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

func (repository *ticketRepositoryES) Insert(idTicket string, ticket *domain.TicketES) error {
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

	return nil
}

func (repository *ticketRepositoryES) Update(idTicket string, ticket *domain.TicketES) error {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	data, err := json.Marshal(ticket)

	if err != nil {
		return err
	}

	req := esapi.UpdateRequest{
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

	return nil
}

