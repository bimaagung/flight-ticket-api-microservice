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

type ESTicketRes struct {
    Index    string 	`json:"_index"`
    ID       string 	`json:"_id"`
    Score    float32  	`json:"_score"`
    Source   Source 	`json:"_source"`
}

type Source struct {
    ID                string 			`json:"id"`
    Track             *domain.Track  	`json:"track"`
    Airplane          *domain.Airplane 	`json:"airplane"`
    ArrivalDatetime   time.Time 			`json:"arrival_datetime"`
    DepartureDatetime time.Time 			`json:"departure_datetime"`
    Price             int    			`json:"price"`
    CreatedAt         time.Time 			`json:"created_at"`
    UpdatedAt         time.Time 			`json:"updated_at"`
}


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

func (repository *ticketRepositoryES) Search(payloadSearch string) ([]*domain.TicketRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": payloadSearch,
				"analyzer": "my_analyzer",
				"fuzziness": "AUTO",
				"fields": []string{"track.arrival", "track.departure"},
				"operator": "or",
				"type": "most_fields",
      			"tie_breaker": 0.3,
			},
		},
	}

	if err:= json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := repository.ESClient.Search(
		repository.ESClient.Search.WithContext(ctx),
		repository.ESClient.Search.WithIndex("ticket"),
		repository.ESClient.Search.WithBody(&buf),
		repository.ESClient.Search.WithTrackTotalHits(true),
		repository.ESClient.Search.WithPretty(),
	)

	 defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	var response struct {
        Hits struct { 
			Total struct { 
				Value int `json:"value"` 
			} `json:"total"` 
		Hits []ESTicketRes `json:"hits"`
		} `json:"hits"`
    }

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	var ticketsRes []*domain.TicketRes

    for _, hit := range response.Hits.Hits {
		trackRes := domain.TrackRes{
			Id: hit.Source.Track.Id.String(),
			Arrival: hit.Source.Track.Arrival,
			Departure: hit.Source.Track.Departure,
		}

		airplaneRes := domain.AirplaneRes{
			Id: hit.Source.Airplane.Id.String(),
			FlightCode: hit.Source.Airplane.FlightCode,
			Seats: hit.Source.Airplane.Seats,
		}

		ticketRes := &domain.TicketRes{
			Id: hit.Source.ID,
			Track: &trackRes,
			Airplane: &airplaneRes,
			ArrivalDatetime: hit.Source.ArrivalDatetime,
			DepartureDatetime: hit.Source.DepartureDatetime,
			Price: hit.Source.Price,
			CreatedAt: hit.Source.CreatedAt,
			UpdatedAt: hit.Source.UpdatedAt,
		}


		ticketsRes = append(ticketsRes, ticketRes)
    }

	return ticketsRes, nil
}

