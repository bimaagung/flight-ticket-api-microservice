package mysqlrepository

import (
	"airplane-service/domain"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

const dbTimeout = time.Second * 3

func NewAirplaneRepositoryMysql(database *sql.DB, rabbitMQ *amqp.Connection) domain.AirplaneRepositoryMysql {
	return &airplaneRepositoryMysql{
		DB: database,
		Rabbit: rabbitMQ,
	}
}

type airplaneRepositoryMysql struct {
	DB *sql.DB
	Rabbit *amqp.Connection
}

func (repository *airplaneRepositoryMysql) Insert(airplane *domain.Airplane)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var ID string = uuid.New().String()
	query := `insert into airplanes (id, flight_code, seats, type, production, factory) values ($1, $2, $3, $4, $5, $6) returning id`

	err := repository.DB.QueryRowContext(ctx, query,
		ID,
		airplane.FlightCode,
		airplane.Seats,
		airplane.Type,
		airplane.Production,
		airplane.Factory,
	).Scan(&ID)

	if err != nil {
		return "", nil
	}

	return ID, nil
}

func (repository *airplaneRepositoryMysql) CheckAirplaneExist(flightCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var airplane domain.Airplane

	query := `select id from airplanes where flight_code = $1 and deleted_at is null`

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

func (repository *airplaneRepositoryMysql) GetById(idAirplane string)(*domain.Airplane, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	airplane := &domain.Airplane{}
	
	query := `select id, flight_code, seats, type, production, factory, created_at, updated_at from airplanes where id = $1 and deleted_at is null`

	err := repository.DB.QueryRowContext(ctx, query, idAirplane).Scan(&airplane.Id, &airplane.FlightCode, &airplane.Seats, &airplane.Type, &airplane.Production, &airplane.Factory, &airplane.CreatedAt, &airplane.UpdatedAt)
	
	if err != nil {
		if err != sql.ErrNoRows{
			return nil, errors.New("airplane not found")
		}

		return nil, err
	}

	return airplane, nil
}

func (repository *airplaneRepositoryMysql) List()([]*domain.Airplane, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `select id, flight_code, seats, type, production, factory, created_at, updated_at from airplanes where deleted_at is null`

	rows, err := repository.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var airplanes []*domain.Airplane

	for rows.Next() {
		var airplane domain.Airplane
		
		err := rows.Scan(
			&airplane.Id, 
			&airplane.FlightCode, 
			&airplane.Seats, 
			&airplane.Type, 
			&airplane.Production, 
			&airplane.Factory, 
			&airplane.CreatedAt, 
			&airplane.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		airplanes = append(airplanes, &airplane)
	}

	return airplanes, nil
}

func (repository *airplaneRepositoryMysql) Delete(idAirplane string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	deletedAt := time.Now()
	
	query := `update airplanes set deleted_at = $1 where id = $2`

	_, err := repository.DB.QueryContext(ctx, query, deletedAt,idAirplane)
	
	if err != nil {
		return err
	}

	return nil
}

func (repository *airplaneRepositoryMysql) VerifyAirplaneAvailable(idAirplane string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `select * from airplanes where id = $1 and deleted_at is null`

	_, err := repository.DB.QueryContext(ctx, query, idAirplane)
	
	if err != nil {
		if err != sql.ErrNoRows{
			return errors.New("airplane not found")
		}

		return err
	}

	return nil
}

func (repository *airplaneRepositoryMysql) Update(idAirplane string, airplane *domain.Airplane)(*domain.Airplane, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	airplaneRes := &domain.Airplane{}
	
	upatedAt := time.Now()

	query := `update airplanes set arrival = $1, departure = $2, long_flight = $3, updated_at = $4 where id = $5 returning id, arrival, departure, long_flight, created_at, updated_at`

	err := repository.DB.QueryRowContext(ctx, query,
		airplane.FlightCode,
		airplane.Seats,
		airplane.Type,
		airplane.Production,
		airplane.Factory,
		upatedAt,
		idAirplane,
	).Scan(
		&airplaneRes.Id,
		airplane.FlightCode,
		airplane.Seats,
		airplane.Type,
		airplane.Production,
		airplane.Factory,
		&airplaneRes.CreatedAt,
		&airplaneRes.UpdatedAt,
	)

	if err != nil {
		return nil, nil
	}

	return airplaneRes, nil
}