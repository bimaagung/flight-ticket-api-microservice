package domain

import (
	"time"

	"github.com/google/uuid"
)

type Track struct {
	Id            	uuid.UUID   `json:"id"`
	Arrival       	string    	`json:"arrival,omitempty"`
	Departure     	string    	`json:"departure,omitempty"`
	LongFlight    	int 		`json:"long_flight,omitempty"`
	CreatedAt 		time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 	`json:"updated_at,omitempty"`
	DeleteAt 		bool 		`json:"deleted_at,omitempty"`
}

type TrackReq struct {
	Arrival       	string    	`json:"arrival,omitempty"`
	Departure     	string    	`json:"departure,omitempty"`
	LongFlight    	int 		`json:"long_flight,omitempty"`
}

type TrackRes struct {
	Id            	string    	`json:"id"`
	Arrival       	string    	`json:"arrival,omitempty"`
	Departure     	string    	`json:"departure,omitempty"`
	LongFlight    	int 	`json:"long_flight,omitempty"`
	CreatedAt 		time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 	`json:"updated_at,omitempty"`
}

type TrackRepositoryPostgres interface {
	Insert(track *Track)(string, error)
	CheckTrackExist(arrival string, depature string) error
	GetById(id string)(*Track, error)
	List()([]*Track, error)
	Delete(id string) error
	VerifyTrackAvailable(idTrack string) error
	Update(idTrack string, track *Track)(*Track, error)
}

type TrackUseCase interface {
	Add(payload *TrackReq)(string, error)
	GetById(idTrack string)(*TrackRes, error)
	List()([]*TrackRes, error)
	Delete(id string) error
	Update(idTrack string, track *TrackReq)(*TrackRes, error)
}