package postgresrepository_test

import (
	"regexp"
	"testing"
	"ticket-service/domain"
	"time"

	postgresrepository "ticket-service/internal/repository/postgres_repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTicketRepositoryPostgres_Insert(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf(err.Error())
	}

	var id uuid.UUID = uuid.New()
	var trackId uuid.UUID = uuid.New()
	var airplaneId uuid.UUID = uuid.New()

	var mockTicket = domain.Ticket{
			Id: id, TrackId: trackId, AirplaneId: airplaneId, Datetime: time.Now(), Price: 1000000,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(mockTicket.Id)

	defer db.Close()
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(
		`insert into tickets (id, track_id, airplane_id, datetime, price) values ($1, $2, $3, $4, $5) returning id`)).
		WithArgs(mockTicket.Id, mockTicket.TrackId, mockTicket.AirplaneId, mockTicket.Datetime, mockTicket.Price).
		WillReturnRows(rows)
	

		ticketRepositoryPostgres := postgresrepository.NewTicketPostgresRepository(db)
		insertTicket, err := ticketRepositoryPostgres.Insert(&mockTicket)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %s", err)
		}

		assert.Equal(t, mockTicket.Id.String(), insertTicket)
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(
		`insert into tickets (id, track_id, airplane_id, datetime, price) values ($1, $2, $3, $4) returning id`)).
		WithArgs(mockTicket.Id, mockTicket.TrackId, mockTicket.AirplaneId, mockTicket.Datetime, mockTicket.Price).
		WillReturnRows(rows)
	

		ticketRepositoryPostgres := postgresrepository.NewTicketPostgresRepository(db)
		_, err := ticketRepositoryPostgres.Insert(&mockTicket)

		assert.Error(t, err)
	})
}