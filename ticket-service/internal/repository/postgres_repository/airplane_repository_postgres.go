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
