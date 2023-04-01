package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"track-service/domain"
	"track-service/helper/event"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

const dbTimeout = time.Second * 3

func NewTrackRepositoryPostgres(database *sql.DB, rabbitMQ *amqp.Connection) domain.TrackRepositoryPostgres {
	return &trackRepositoryPostgres{
		DB: database,
		Rabbit: rabbitMQ,
	}
}

type trackRepositoryPostgres struct {
	DB *sql.DB
	Rabbit *amqp.Connection
}

func (repository *trackRepositoryPostgres) Insert(track *domain.Track)(string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
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

	// push to queue
	emitter, err := event.NewEventEmitter(repository.Rabbit)
	
	if err != nil {
		return "", err
	}

	err = emitter.PushToQueue(&domain.TrackRes{
		Id: ID.String(),
		Arrival: track.Arrival,
		Departure: track.Departure,
		LongFlight: track.LongFlight,
	}, "track.INFO")
	
	if err != nil {
		return "", err
	}

	return ID.String(), nil
}

func (repository *trackRepositoryPostgres) CheckTrackExist(arrival string, departure string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
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

func (repository *trackRepositoryPostgres) GetById(idTrack string)(*domain.Track, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	track := &domain.Track{}

	uuidConvert, err := uuid.Parse(idTrack)

	if err != nil {
		return nil, errors.New("track not found")
	}
	
	query := `select id, arrival, departure, long_flight, created_at from tracks where id = $1 and deleted_at is null`

	err = repository.DB.QueryRowContext(ctx, query, uuidConvert).Scan(&track.Id, &track.Arrival, &track.Departure, &track.LongFlight, &track.CreatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errors.New("track not found")
		}

		return nil, err
	}

	return track, nil
}

func (repository *trackRepositoryPostgres) List()([]*domain.Track, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `select id, arrival, departure, long_flight, created_at from tracks where deleted_at is null`

	rows, err := repository.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tracks []*domain.Track

	for rows.Next() {
		var track domain.Track
		
		err := rows.Scan(
			&track.Id,
			&track.Arrival,
			&track.Departure,
			&track.LongFlight,
			&track.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (repository *trackRepositoryPostgres) Delete(idTrack string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	deletedAt := time.Now()
	
	uuidConvert, err := uuid.Parse(idTrack)

	if err != nil {
		return errors.New("track not found")
	}
	
	query := `update tracks set deleted_at = $1 where id = $2`

	_, err = repository.DB.QueryContext(ctx, query, deletedAt,uuidConvert)
	
	if err != nil {
		return err
	}

	return nil
}

func (repository *trackRepositoryPostgres) VerifyTrackAvailable(idTrack string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	uuidConvert, err := uuid.Parse(idTrack)

	if err != nil {
		return err
	}

	var Id uuid.UUID
	
	query := `select id from tracks where id = $1 and deleted_at is null`

	err = repository.DB.QueryRowContext(ctx, query, uuidConvert).Scan(&Id)
	
	if err != nil {
		if err == sql.ErrNoRows{
			return errors.New("track not found")
		}

		return err
	}

	return nil
}

func (repository *trackRepositoryPostgres) Update(idTrack string, track *domain.Track)(*domain.Track, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	trackRes := &domain.Track{}
	upatedAt := time.Now()

	uuidConvert, err := uuid.Parse(idTrack)

	if err != nil {
		return nil, errors.New("track not found")
	}

	query := `update tracks set arrival = $1, departure = $2, long_flight = $3, updated_at = $4 where id = $5 returning id, arrival, departure, long_flight, created_at, updated_at`

	err = repository.DB.QueryRowContext(ctx, query,
		track.Arrival,
		track.Departure,
		track.LongFlight,
		upatedAt,
		uuidConvert,
	).Scan(
		&trackRes.Id,
		&trackRes.Arrival,
		&trackRes.Departure,
		&trackRes.LongFlight,
		&trackRes.CreatedAt,
		&trackRes.UpdatedAt,
	)

	if err != nil {
		return nil, nil
	}

	return trackRes, nil
}