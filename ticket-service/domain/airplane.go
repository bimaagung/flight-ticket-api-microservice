package domain

import (
	"time"

	"github.com/google/uuid"
)

type Airplane struct {
	Id          uuid.UUID 	`json:"id"`
	FlightCode  string 		`json:"flight_code,omitempty"`
	Seats       int    		`json:"seats,omitempty"`
	CreatedAt 	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`json:"updated_at,omitempty"`
	DeleteAt 	bool 		`json:"deleted_at,omitempty"`
}

type AirplaneRes struct {
	Id          string 	`json:"id"`
	FlightCode  string 		`json:"flight_code,omitempty"`
	Seats       int    		`json:"seats,omitempty"`
}

type AirplaneUseCase interface {
	Add(payload *Airplane) (string, error)
}

type AirplaneRepositoryPostgres interface {
	VerifyAirplaneAvailable(idTrack string) error
	Insert(airplane *Airplane) (string, error)
	CheckAirplaneExist(flightCode string) error
	List()([]*Airplane, error)
}