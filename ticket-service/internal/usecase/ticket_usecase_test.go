package usecase_test

import (
	"testing"
	"ticket-service/domain"
	mockses "ticket-service/internal/mocks/es"
	mockspostgres "ticket-service/internal/mocks/postgres"
	"ticket-service/internal/usecase"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTicketUC_Add(t *testing.T) {

	mockTicketPostgresRepository := new(mockspostgres.TicketPostgresRepository)
	mockTrackPostgresRepository := new(mockspostgres.TrackPostgresRepository)
	mockAirplanePostgresRepository := new(mockspostgres.AirplanePostgresRepository)
	mockTicketESRepository := new(mockses.TicketESRepository)

	

	var idTicket string = "0f10cff9-4b5a-4e5a-b976-56bf56efa280"

	parseId, err := uuid.Parse(idTicket)

	if err != nil {
		t.Errorf("Parse id to uuid error: %s", err)
	}

	var payload *domain.TicketReq = &domain.TicketReq{
		TrackId: "0f10cff9-4b5a-4e5a-b976-56bf56efa280",
		AirplaneId: "0f10cff9-4b5a-4e5a-b976-56bf56efa280",
		Datetime: "2022-12-21T13:00:00Z",
		Price: 1000000,
	} 

	parseTime, err := time.Parse(time.RFC3339, payload.Datetime)

	if err != nil {
		t.Errorf("Parse time error: %s", err)
	}

	var ticket *domain.Ticket = &domain.Ticket{
		Id: parseId,
		TrackId: parseId,
		AirplaneId: parseId,
		Datetime: parseTime,
		Price: payload.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var track *domain.Track = &domain.Track{
		Id: parseId,
		Arrival: "Jakarta",
		Departure: "Semarang",
		LongFlight: 40,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var airplane *domain.Airplane = &domain.Airplane{
		Id: parseId,
		FlightCode: "GH098J",
		Seats: 230,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}  

	t.Run("success", func(t *testing.T) {
		mockTicketPostgresRepository.On("CheckTicketExist", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("time.Time")).Return(nil).Once()
		mockTrackPostgresRepository.On("VerifyTrackAvailable", mock.AnythingOfType("string")).Return(nil).Once()
		mockAirplanePostgresRepository.On("VerifyAirplaneAvailable", mock.AnythingOfType("string")).Return(nil).Once()
		mockTicketPostgresRepository.On("Insert", mock.Anything).Return(idTicket,nil).Once()
		mockTicketPostgresRepository.On("GetById", mock.AnythingOfType("string")).Return(ticket, track, airplane, nil).Once()
		mockTicketESRepository.On("Insert", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

		uc := usecase.NewTicketUseCase(
			mockTicketPostgresRepository, 
			mockTrackPostgresRepository, 
			mockAirplanePostgresRepository, 
			mockTicketESRepository,
		)

		id, err := uc.Add(payload)

		assert.NoError(t, err)
		assert.NotNil(t, id)

		mockTicketPostgresRepository.AssertExpectations(t)
		mockTrackPostgresRepository.AssertExpectations(t)
		mockAirplanePostgresRepository.AssertExpectations(t)
		mockTicketESRepository.AssertExpectations(t)
	})
}