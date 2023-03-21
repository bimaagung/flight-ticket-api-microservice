package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
)


func NewTrackRepositoryPostgres(database *sql.DB) domain.TrackRepositoryPostgres {
	return &trackRepositoryPostgres{
		DB: database,
		DBTimeout: time.Second * 3,
	}
}

type trackRepositoryPostgres struct {
	DB *sql.DB
	DBTimeout time.Duration
}

func (repository *trackRepositoryPostgres) CheckTrackExist(arrival string, departure string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var track domain.Track
	query := `select id from tracks where arrival = $1 and departure = $2 and deleted_at is null`

	row := repository.DB.QueryRowContext(ctx, query, arrival, departure)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(
		&track.Id,
	)

	if err == sql.ErrNoRows{
		return nil
	}

	if err != sql.ErrNoRows{
		return errors.New("track is exist")
	}

	return err

}

func (repository *trackRepositoryPostgres) VerifyTrackAvailable(idTrack string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	uuidConvert, err := uuid.Parse(idTrack)

	if err != nil {
		return errors.New("track not found")
	}
	
	query := `select * from tracks where id = $1 and deleted_at is null`

	_, err = repository.DB.QueryContext(ctx, query, uuidConvert)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return errors.New("track not found")
		}

		return err
	}

	return nil
}

func (repository *trackRepositoryPostgres) Insert(track *domain.Track)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repository.DBTimeout)
	defer cancel()

	var ID uuid.UUID = uuid.New()
	query := `insert into tracks (id, arrival, departure, long_flight) values ($1, $2, $3, $4) returning id`

	err := repository.DB.QueryRowContext(ctx, query,
		ID,
		track.Arrival,
		track.Departure,
		track.LongFlight,
	).Scan(&ID)

	if err != nil {
		return "", err
	}

	return ID.String(), nil
}
