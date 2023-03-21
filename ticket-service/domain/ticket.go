package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id            	uuid.UUID   	`json:"id"`
	TrackId       	uuid.UUID    	`json:"track_id,omitempty"`
	AirplaneId     	uuid.UUID    	`json:"airplane_id,omitempty"`
	Date    		time.Time 		`json:"date,omitempty"`
	Time    		time.Time 		`json:"time,omitempty"`
	Price    		int 			`json:"price,omitempty"`
	CreatedAt 		time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 		`json:"updated_at,omitempty"`
	DeletedAt 		time.Time 		`json:"deleted_at,omitempty"`
}

type TicketReq struct {
	TrackId       	string    	`json:"track_id,omitempty"`
	AirplaneId     	string    	`json:"airplane_id,omitempty"`
	Date    		time.Time 		`json:"date,omitempty"`
	Time    		time.Time 		`json:"time,omitempty"`
	Price    		int 			`json:"price,omitempty"`
}

type TicketRes struct {
	Id            	uuid.UUID   	`json:"id"`
	Track       	TrackRes    	`json:"track,omitempty"`
	AirplaneId     	AirplaneRes    	`json:"airplane,omitempty"`
	Date    		int 			`json:"date,omitempty"`
	Time    		int 			`json:"time,omitempty"`
	Price    		int 			`json:"price,omitempty"`
	CreatedAt 		time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 		`json:"updated_at,omitempty"`
	DeleteAt 		bool 			`json:"deleted_at,omitempty"`
}

type TicketRepositoryPostgres interface {
	Insert(ticket *Ticket)(string, error)
	CheckTicketExist(trackId uuid.UUID, airplaneId uuid.UUID, date time.Time, time time.Time) error
	Delete(id string) error
	VerifyTicketAvailable(idTicket string) error
	Update(idTicket string, ticket *Ticket) error
}

type TicketUseCase interface {
	Add(payload *TicketReq)(string, error)
	Delete(id string) error
	Update(idTicket string, ticket *TicketReq) error
}