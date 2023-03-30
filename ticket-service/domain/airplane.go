package domain

import (
	"github.com/google/uuid"
)

type Airplane struct {
	Id            		uuid.UUID   	`json:"id"`
	FlightCode       	string    		`json:"flight_code,omitempty"`
	Seats    			int    			`json:"seats,omitempty"`
}

type AirplaneUseCase interface {
	Add(payload *Airplane)(string, error)
}

type AirplaneRepositoryPostgres interface {
	VerifyAirplaneAvailable(idTrack string) error
	Insert(airplane *Airplane)(string, error)
	CheckAirplaneExist(flightCode string) error
}