package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
)

const dbTimeout = time.Second * 3

func NewTicketPostgresRepository(database *sql.DB) domain.TicketRepositoryPostgres {
	return &ticketRepositoryPostgres{
		DB: database,
	}	
}

type ticketRepositoryPostgres struct {
	DB *sql.DB
}

func(repository *ticketRepositoryPostgres) Insert(ticket *domain.Ticket)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	var ID uuid.UUID = uuid.New()
	query := `insert into ticket (id, ticket_id, airplane_id, date, time, price) VALUES ($1, $2, $3, $4, $5, $6)`

	err := repository.DB.QueryRowContext(ctx, query, 
		ID,
		ticket.TrackId,
		ticket.AirplaneId,
		ticket.Date,
		ticket.Time,
		ticket.Price,
	).Scan(&ID)

	if err != nil {
		return "", err
	}

	return ID.String(), nil
}

func (repository *ticketRepositoryPostgres) CheckTicketExist(trackId uuid.UUID, airplaneId uuid.UUID, date time.Time, time time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var ticket domain.Ticket
	query := `select id from tickets where track_id = $1 and airplane_id = $2 and date = $3 and time = $4 and deleted_at is null`

	row := repository.DB.QueryRowContext(ctx, query, trackId, airplaneId, date, time)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(
		&ticket.Id,
	)

	if err == sql.ErrNoRows{
		return nil
	}

	if err != sql.ErrNoRows{
		return errors.New("ticket is exist")
	}

	return err

}

func (repository *ticketRepositoryPostgres) Delete(idTicket string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	deletedAt := time.Now()
	
	uuidConvert, err := uuid.Parse(idTicket)

	if err != nil {
		return errors.New("ticket not found")
	}
	
	query := `update tickets set deleted_at = $1 where id = $2`

	_, err = repository.DB.QueryContext(ctx, query, deletedAt, uuidConvert)
	
	if err != nil {
		return err
	}

	return nil
}

func (repository *ticketRepositoryPostgres) VerifyTicketAvailable(idTicket string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	uuidConvert, err := uuid.Parse(idTicket)

	if err != nil {
		return errors.New("ticket not found")
	}
	
	query := `select * from tickets where id = $1 and deleted_at is null`

	_, err = repository.DB.QueryContext(ctx, query, uuidConvert)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return errors.New("ticket not found")
		}

		return err
	}

	return nil
}

func (repository *ticketRepositoryPostgres) Update(idTicket string, ticket *domain.Ticket)error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	upatedAt := time.Now()

	uuidConvert, err := uuid.Parse(idTicket)

	if err != nil {
		return errors.New("ticket not found")
	}

	query := `update tickets set track_id = $1, airplane_id = $2, date = $3, time = $4, price = $5, updated_at = $7 where id = $5`

	result, err := repository.DB.ExecContext(ctx, query, 
		ticket.TrackId,
		ticket.AirplaneId,
		ticket.Date,
		ticket.Time,
		ticket.Price,
		upatedAt,
		uuidConvert,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("can't update ticket ")
	}

	return nil
}