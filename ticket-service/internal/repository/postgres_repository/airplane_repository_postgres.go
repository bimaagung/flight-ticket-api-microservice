package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
)


func NewAirplaneRepositoryPostgres(database *sql.DB) domain.AirplaneRepositoryPostgres {
	return &airplaneRepositoryPostgres{
		DB: database,
		DBTimeout: time.Second * 3,
	}
}

type airplaneRepositoryPostgres struct {
	DB *sql.DB
	DBTimeout time.Duration
}

func (repository *airplaneRepositoryPostgres) VerifyAirplaneAvailable(idAirplane string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	uuidConvert, err := uuid.Parse(idAirplane)

	if err != nil {
		return errors.New("airplane not found")
	}
	
	query := `select * from airplanes where id = $1 and deleted_at is null`

	_, err = repository.DB.QueryContext(ctx, query, uuidConvert)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return errors.New("airplane not found")
		}

		return err
	}

	return nil
}

func (repository *airplaneRepositoryPostgres) Insert(airplane *domain.Airplane)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var ID uuid.UUID = uuid.MustParse(airplane.Id.String())
	query := `insert into airplanes (id, flight_code, seats) values ($1, $2, $3) returning id`

	err := repository.DB.QueryRowContext(ctx, query,
		ID,
		airplane.FlightCode,
		airplane.Seats,
	).Scan(&ID)

	if err != nil {
		return "", err
	}

	return ID.String(), nil
}


func (repository *airplaneRepositoryPostgres) CheckAirplaneExist(flightCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var airplane domain.Airplane
	query := `select id from airplanes where flight_code = $1`

	row := repository.DB.QueryRowContext(ctx, query, flightCode)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(
		&airplane.Id,
	)

	if err == sql.ErrNoRows{
		return nil
	}

	if err != sql.ErrNoRows{
		return errors.New("airplane is exist")
	}

	return err

}