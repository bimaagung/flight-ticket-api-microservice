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
		ticketRepositoryES domain.TicketRepositoryElasticsearch,
	) domain.TicketUseCase {

	return &ticketUseCaseImpl{
		TicketRepositoryPostgres: ticketRepositoryPostgres,
		TrackRepositoryPostgres: trackRepositoryPostgres,
		AirplaneRepositoryPostgres: airplaneRepositoryPostgres,
		ticketRepositoryES: ticketRepositoryES,
	}

}

type ticketUseCaseImpl struct {
	TicketRepositoryPostgres domain.TicketRepositoryPostgres
	TrackRepositoryPostgres domain.TrackRepositoryPostgres
	AirplaneRepositoryPostgres domain.AirplaneRepositoryPostgres
	ticketRepositoryES domain.TicketRepositoryElasticsearch
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

	err = useCase.TicketRepositoryPostgres.CheckTicketExist(parseTrackId, parseAirplalneId, parseTime)

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

	err = useCase.ticketRepositoryES.Insert(id, &domain.TicketES{
		TrackId: parseTrackId,
		AirplaneId: parseAirplalneId,
		Datetime: parseTime,
		Price: payload.Price,
	})

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

	var ticketRes []*domain.TicketRes

	tickets, err := useCase.TicketRepositoryPostgres.List()

	if err != nil {
		return nil, err
	}

	airplanes, err := useCase.AirplaneRepositoryPostgres.List()

	if err != nil {
		return nil, err
	}

	tracks, err := useCase.TrackRepositoryPostgres.List()

	if err != nil {
		return nil, err
	}


	for _, v := range tickets {

		var longFlight int
		track := &domain.TrackRes{}
		airplane := &domain.AirplaneRes{}

		for _, t := range tracks {
			if t.Id == v.TrackId {
				track.Id = t.Id.String()
				track.Arrival = t.Arrival
				track.Departure = t.Departure
				track.LongFlight = t.LongFlight
				longFlight = t.LongFlight
			}
		}

		for _, a := range airplanes {
			if a.Id == v.AirplaneId {
				airplane.Id = a.Id.String()
				airplane.FlightCode = a.FlightCode
				airplane.Seats = a.Seats
			}
		}


		durasi := time.Duration(longFlight) * time.Minute
		arrivalDatetime := v.Datetime.Add(durasi) 

		result := &domain.TicketRes{
			Id: v.Id.String(),
			Track: track,
			Airplane: airplane,
			ArrivalDatetime: arrivalDatetime,
			DepartureDatetime: v.Datetime,
			Price: v.Price,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		ticketRes = append(ticketRes, result)
	}

	return ticketRes, nil
}