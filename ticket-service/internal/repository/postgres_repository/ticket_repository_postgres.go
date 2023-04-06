package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
)

func NewTicketPostgresRepository(database *sql.DB) domain.TicketRepositoryPostgres {
	return &ticketRepositoryPostgres{
		DB: database,
		DBTimeout: time.Second * 3,
	}	
}

type ticketRepositoryPostgres struct {
	DB *sql.DB
	DBTimeout time.Duration
}

func(repository *ticketRepositoryPostgres) Insert(ticket *domain.Ticket)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),repository.DBTimeout)
	defer cancel()

	var id uuid.UUID = uuid.New()

	query := `insert into tickets (id, track_id, airplane_id, datetime, price) VALUES ($1, $2, $3, $4, $5) returning id`

	err := repository.DB.QueryRowContext(ctx, query, 
		id,
		ticket.TrackId,
		ticket.AirplaneId,
		ticket.Datetime,
		ticket.Price,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (repository *ticketRepositoryPostgres) CheckTicketExist(trackId uuid.UUID, airplaneId uuid.UUID, datetime time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var id uuid.UUID
	query := `select id from tickets where track_id = $1 and airplane_id = $2 and datetime = $3 and deleted_at is null`

	row := repository.DB.QueryRowContext(ctx, query, trackId, airplaneId, datetime)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(
		&id,
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
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	deletedAt := time.Now()

	parseId, err := uuid.Parse(idTicket)

	if err != nil {
		return errors.New("track not found")
	}

	query := `update tickets set deleted_at = $1 where id = $2`

	_, err = repository.DB.QueryContext(ctx, query, deletedAt, parseId)
	
	if err != nil {
		return err
	}

	return nil
}

func (repository *ticketRepositoryPostgres) VerifyTicketAvailable(idTicket string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()
	
	var id uuid.UUID

	query := `select id from tickets where id = $1 and deleted_at is null`

	err := repository.DB.QueryRowContext(ctx, query, idTicket).Scan(&id)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return errors.New("ticket not found")
		}

		return err
	}

	return nil
}

func (repository *ticketRepositoryPostgres) Update(idTicket string, ticket *domain.Ticket)error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	upatedAt := time.Now()

	parseId, err := uuid.Parse(idTicket)

	if err != nil {
		return err
	}

	query := `update tickets set track_id = $1, airplane_id = $2, datetime = $3, price = $4, updated_at = $5 where id = $6 and deleted_at is null`

	result, err := repository.DB.ExecContext(ctx, query, 
		ticket.TrackId,
		ticket.AirplaneId,
		ticket.Datetime,
		ticket.Price,
		upatedAt,
		parseId,
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

func (repository *ticketRepositoryPostgres) GetById(idTicket string)(*domain.Ticket, *domain.Track, *domain.Airplane, error){
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	ticket := &domain.Ticket{}
	track := &domain.Track{}
	airplane := &domain.Airplane{}
	
	
	parseId, err := uuid.Parse(idTicket)
	if err != nil {
		return nil, nil, nil, err
	}

	query := `select t.id, t.datetime, t.price, tr.id, tr.arrival, tr.departure, tr.long_flight, a.id, a.flight_code, a.seats, t.created_at, t.updated_at 
			from tickets t
			inner join airplanes a on t.airplane_id = a.id 
			inner join tracks tr on t.track_id = tr.id 
			where t.id = $1 and t.deleted_at is null`

	err = repository.DB.QueryRowContext(ctx, query, parseId).Scan(
		&ticket.Id, 
		&ticket.Datetime, 
		&ticket.Price, 
		&track.Id, 
		&track.Arrival, 
		&track.Departure,
		&track.LongFlight,
		&airplane.Id, 
		&airplane.FlightCode, 
		&airplane.Seats, 
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil, nil, errors.New("ticket not found")
		}

		return nil, nil, nil, err
	}

	return ticket, track, airplane, nil
}

func (repository *ticketRepositoryPostgres) List()([]*domain.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var result []*domain.Ticket

	query := `select id, track_id, airplane_id ,datetime, price, created_at, updated_at from tickets where deleted_at is null order by created_at desc`
	
	rows, err := repository.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var ticket domain.Ticket

		err := rows.Scan(
			&ticket.Id, 
			&ticket.TrackId, 
			&ticket.AirplaneId, 
			&ticket.Datetime, 
			&ticket.Price, 
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}


		result = append(result, &ticket)

	}

	return result, nil
	
}