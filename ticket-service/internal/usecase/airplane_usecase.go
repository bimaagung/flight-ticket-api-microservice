package usecase

import (
	"ticket-service/domain"
)

func NewAirplaneUseCase(airplaneRepositoryPostgres domain.AirplaneRepositoryPostgres) domain.AirplaneUseCase {
	return &airplaneUseCaseImpl{
		AirplaneRepositoryPostgres: airplaneRepositoryPostgres,
	}
}

type airplaneUseCaseImpl struct {
	AirplaneRepositoryPostgres domain.AirplaneRepositoryPostgres
}

func(useCase *airplaneUseCaseImpl) Add(payload *domain.Airplane)(string, error){
	
	err := useCase.AirplaneRepositoryPostgres.CheckAirplaneExist(payload.FlightCode)

	if err != nil {
		return "", err
	}

	id, err := useCase.AirplaneRepositoryPostgres.Insert(payload)

	if err != nil {
		return "", err
	}

	return id, nil
}
