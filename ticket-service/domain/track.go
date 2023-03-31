package domain

import (
	"time"

	"github.com/google/uuid"
)

type Track struct {
	Id            	uuid.UUID  	`json:"id"`
	Arrival       	string    	`json:"arrival,omitempty"`
	Departure     	string    	`json:"departure,omitempty"`
	LongFlight    	int 		`json:"long_flight,omitempty"`
	CreatedAt 		time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 	`json:"updated_at,omitempty"`
	DeleteAt 		bool 		`json:"deleted_at,omitempty"`
}

type TrackRes struct {
	Id         string    `json:"id"`
	Arrival    string    `json:"arrival,omitempty"`
	Departure  string    `json:"departure,omitempty"`
	LongFlight int       `json:"long_flight,omitempty"`
}

type TrackRepositoryPostgres interface {
	CheckTrackExist(arrival string, depature string) error
	VerifyTrackAvailable(idTrack string) error
	Insert(track *Track)(string, error)
}

type TrackUseCase interface {
	Add(payload *Track)(string, error)
}