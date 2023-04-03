package usecase

import (
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
)

func NewTicketUseCase(
		ticketRepositoryPostgres domain.TicketRepositoryPostgres,
		trackRepositoryPostgres domain.TrackRepositoryPostgres,
		airplaneRepositoryPostgres domain.AirplaneRepositoryPostgres,
	) domain.TicketUseCase {

	return &ticketUseCaseImpl{
		TicketRepositoryPostgres: ticketRepositoryPostgres,
		TrackRepositoryPostgres: trackRepositoryPostgres,
		AirplaneRepositoryPostgres: airplaneRepositoryPostgres,
	}

}

type ticketUseCaseImpl struct {
	TicketRepositoryPostgres domain.TicketRepositoryPostgres
	TrackRepositoryPostgres domain.TrackRepositoryPostgres
	AirplaneRepositoryPostgres domain.AirplaneRepositoryPostgres
}

func(useCase *ticketUseCaseImpl) Add(payload *domain.TicketReq)(string, error){

	parseTime, err := time.Parse(time.RFC3339, payload.Datetime)

	if err != nil {
		return "", err
	}

	parseTrackId, err := uuid.Parse(payload.TrackId)

	if err != nil {
		return "", err
	}

	parseAirplalneId, err := uuid.Parse(payload.AirplaneId)

	if err != nil {
		return "", err
	}

	ticket := &domain.Ticket{
		TrackId: parseTrackId,
		AirplaneId: parseAirplalneId,
		Datetime: parseTime,
		Price: payload.Price,
	}

	err = useCase.TicketRepositoryPostgres.CheckTicketExist(payload.TrackId, payload.AirplaneId, ticket.Datetime)

	if err != nil {
		return "", err
	}

	// check if track not found
	err = useCase.TrackRepositoryPostgres.VerifyTrackAvailable(payload.TrackId)
	if err != nil {
		return "", err
	}

	// check if airplane not found
	err = useCase.AirplaneRepositoryPostgres.VerifyAirplaneAvailable(payload.AirplaneId)
	if err != nil {
		return "", err
	}

	id, err := useCase.TicketRepositoryPostgres.Insert(ticket)

	if err != nil {
		return "", err
	}

	return id, nil
}

func(useCase *ticketUseCaseImpl) Delete(id string) error {

	err := useCase.TicketRepositoryPostgres.VerifyTicketAvailable(id)

	if err != nil {
		return err
	}
	
	err = useCase.TicketRepositoryPostgres.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func(useCase *ticketUseCaseImpl) Update(idTicket string, payload *domain.TicketReq) error {

	parseTime, err := time.Parse(time.RFC3339, payload.Datetime)

	if err != nil {
		return err
	}

	parseTrackId, err := uuid.Parse(payload.TrackId)

	if err != nil {
		return err
	}

	parseAirplalneId, err := uuid.Parse(payload.AirplaneId)

	if err != nil {
		return err
	}

	ticket := &domain.Ticket{
		TrackId: parseTrackId,
		AirplaneId: parseAirplalneId,
		Datetime: parseTime,
		Price: payload.Price,
	}

	err = useCase.TicketRepositoryPostgres.VerifyTicketAvailable(idTicket)

	if err != nil {
		return err
	}

	// check if track not found
	err = useCase.TrackRepositoryPostgres.VerifyTrackAvailable(payload.TrackId)
	
	if err != nil {
		return err
	}

	// check if airplane not found
	err = useCase.AirplaneRepositoryPostgres.VerifyAirplaneAvailable(payload.AirplaneId)
	
	if err != nil {
		return err
	}

	err = useCase.TicketRepositoryPostgres.Update(idTicket, ticket)

	if err != nil {
		return err
	}

	return nil
}

func(useCase *ticketUseCaseImpl) GetById(id string)(*domain.TicketRes, error){

	ticket, track ,airplane, err := useCase.TicketRepositoryPostgres.GetById(id)

	if err != nil {
		return nil, err
	}

	durasi := time.Duration(track.LongFlight) * time.Minute
	arrivalDatetime := ticket.Datetime.Add(durasi) 

	trackRes := &domain.TrackRes{
		Id: track.Id.String(),
		Arrival: track.Arrival,
		Departure: track.Departure,
		LongFlight: track.LongFlight,
	}

	airplaneRes := &domain.AirplaneRes{
		Id: airplane.Id.String(),
		FlightCode: airplane.FlightCode,
		Seats: airplane.Seats,
	}

	result := &domain.TicketRes{
		Id: 				ticket.Id.String(),
		Track: 				trackRes,
		Airplane: 			airplaneRes,
		ArrivalDatetime: 	arrivalDatetime,
		DepartureDatetime: 	ticket.Datetime,
		Price: 				ticket.Price,
		CreatedAt: 			ticket.CreatedAt,
		UpdatedAt: 			ticket.UpdatedAt,
	} 

	return result, nil
}

func(useCase *ticketUseCaseImpl) List()([]*domain.TicketRes, error) {
	tickets, err := useCase.TicketRepositoryPostgres.List()

	if err != nil {
		return nil, err
	}

	return tickets, nil
}