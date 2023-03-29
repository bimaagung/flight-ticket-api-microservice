package usecase

import (
	"log"
	"ticket-service/domain"
)

func NewTrackUseCase(trackRepositoryPostgres domain.TrackRepositoryPostgres) domain.TrackUseCase {
	return &trackUseCaseImpl{
		TrackRepositoryPostgres: trackRepositoryPostgres,
	}
}

type trackUseCaseImpl struct {
	TrackRepositoryPostgres domain.TrackRepositoryPostgres
}

func(useCase *trackUseCaseImpl) Add(payload *domain.Track)(string, error){
	
	log.Println(payload)
	err := useCase.TrackRepositoryPostgres.CheckTrackExist(payload.Arrival, payload.Departure)

	if err != nil {
		return "", err
	}

	id, err := useCase.TrackRepositoryPostgres.Insert(payload)

	if err != nil {
		return "", err
	}

	return id, nil
}
