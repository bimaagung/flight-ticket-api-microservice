package usecase

import (
	"errors"
	"ticket-service/domain"

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
	
	trackId, err := uuid.Parse(payload.TrackId)

	if err != nil {
		return "", errors.New("can't add ticket")
	}

	airplaneId, err := uuid.Parse(payload.AirplaneId)

	if err != nil {
		return "", errors.New("can't add ticket")
	}

	ticket := &domain.Ticket{
		TrackId: trackId,
		AirplaneId: airplaneId,
		Date: payload.Date,
		Time: payload.Time,
	}

	err = useCase.TicketRepositoryPostgres.CheckTicketExist(trackId, airplaneId, ticket.Date, ticket.Time)

	if err != nil {
		return "", err
	}

	// check if track not found
	err = useCase.TrackRepositoryPostgres.VerifyTrackAvailable(payload.TrackId)
	if err != nil {
		return "", err
	}

	// check if airplane not found
	err = useCase.AirplaneRepositoryPostgres.VerifyAirplaneAvailable(payload.TrackId)
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
	
	trackId, err := uuid.Parse(payload.TrackId)

	if err != nil {
		return errors.New("can't update ticket")
	}

	airplaneId, err := uuid.Parse(payload.AirplaneId)

	if err != nil {
		return errors.New("can't updated ticket")
	}

	ticket := &domain.Ticket{
		TrackId: trackId,
		AirplaneId: airplaneId,
		Date: payload.Date,
		Time: payload.Time,
	}

	err = useCase.TicketRepositoryPostgres.VerifyTicketAvailable(idTicket)

	if err != nil {
		return err
	}

	err = useCase.TicketRepositoryPostgres.Update(idTicket, ticket)

	if err != nil {
		return err
	}

	return nil
}