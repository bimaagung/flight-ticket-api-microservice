package usecase

import (
	"airplane-service/domain"
	"time"
)

const formatTime = "2006-01-02"

func NewAirplaneUseCase(airplaneRepositoryMysql domain.AirplaneRepositoryMysql) domain.AirplaneUseCase {
	return &airplaneUseCaseImpl{
		AirplaneRepositoryMysql: airplaneRepositoryMysql,
	}
}

type airplaneUseCaseImpl struct {
	AirplaneRepositoryMysql domain.AirplaneRepositoryMysql
}

func(useCase *airplaneUseCaseImpl) Add(payload *domain.AirplaneReq)(string, error){
	
	parseProduction, err := time.Parse(formatTime, payload.ProductionDate)
	
	if err != nil {
		return "", err
	}

	airplane := &domain.Airplane{
		FlightCode: payload.FlightCode,
		Seats: payload.Seats,
		Type: payload.Type,
		ProductionDate: parseProduction,
		Factory: payload.Factory,
	}

	err = useCase.AirplaneRepositoryMysql.CheckAirplaneExist(payload.FlightCode)

	if err != nil {
		return "", err
	}

	id, err := useCase.AirplaneRepositoryMysql.Insert(airplane)

	if err != nil {
		return "", err
	}

	return id, nil
}

func(useCase *airplaneUseCaseImpl) GetById(id string)(*domain.AirplaneRes, error){

	airplane, err := useCase.AirplaneRepositoryMysql.GetById(id)

	if err != nil {
		return nil, err
	}

	result := &domain.AirplaneRes{
		Id: airplane.Id,
		FlightCode: airplane.FlightCode,
		Seats: airplane.Seats,
		Type: airplane.Type,
		ProductionDate: airplane.ProductionDate,
		Factory: airplane.Factory,
		CreatedAt: airplane.CreatedAt,
		UpdatedAt: airplane.UpdatedAt,
	}

	return result, nil
}

func(useCase *airplaneUseCaseImpl) List()([]*domain.AirplaneRes, error){

	airplane, err := useCase.AirplaneRepositoryMysql.List()

	if err != nil {
		return nil, err
	}

	var airplanes []*domain.AirplaneRes

	for _, v := range airplane {
		var airplane = domain.AirplaneRes{
			Id: v.Id,
			FlightCode: v.FlightCode,
			Seats: v.Seats,
			Type: v.Type,
			ProductionDate: v.ProductionDate,
			Factory: v.Factory,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		airplanes = append(airplanes, &airplane)
		
	}


	return airplanes, nil
}

func(useCase *airplaneUseCaseImpl) Delete(id string) error {

	err := useCase.AirplaneRepositoryMysql.VerifyAirplaneAvailable(id)
	if err != nil {
		return err
	}
	
	err = useCase.AirplaneRepositoryMysql.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func(useCase *airplaneUseCaseImpl) Update(idAirplane string, payload *domain.AirplaneReq)(*domain.AirplaneRes, error){
	
	parseProduction, err := time.Parse(formatTime, payload.ProductionDate)
	if err != nil {
		return nil, err
	}

	airplane := &domain.Airplane{
		FlightCode: payload.FlightCode,
		Seats: payload.Seats,
		Type: payload.Type,
		ProductionDate: parseProduction,
		Factory: payload.Factory,
	}

	err = useCase.AirplaneRepositoryMysql.VerifyAirplaneAvailable(idAirplane)

	if err != nil {
		return nil, err
	}

	airplaneRes, err := useCase.AirplaneRepositoryMysql.Update(idAirplane, airplane)

	if err != nil {
		return nil, err
	}

	result := &domain.AirplaneRes{
		Id: airplaneRes.Id,
		FlightCode: airplaneRes.FlightCode,
		Seats: airplaneRes.Seats,
		Type: airplaneRes.Type,
		ProductionDate: airplaneRes.ProductionDate,
		Factory: airplaneRes.Factory,
		CreatedAt: airplaneRes.CreatedAt,
		UpdatedAt: airplaneRes.UpdatedAt,
	}

	return result, nil
}