package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id            	uuid.UUID   	`json:"id"`
	TrackId       	uuid.UUID    	`json:"track_id,omitempty"`
	AirplaneId     	uuid.UUID    	`json:"airplane_id,omitempty"`
	Datetime    	time.Time 		`json:"date,omitempty"`
	Price    		int 			`json:"price,omitempty"`
	CreatedAt 		time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 		`json:"updated_at,omitempty"`
	DeletedAt 		time.Time 		`json:"deleted_at,omitempty"`
}

type TicketES struct {
	TrackId       	uuid.UUID    	`json:"track_id,omitempty"`
	AirplaneId     	uuid.UUID    	`json:"airplane_id,omitempty"`
	Datetime    	time.Time 		`json:"date,omitempty"`
	Price    		int 			`json:"price,omitempty"`
	CreatedAt 		time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 		`json:"updated_at,omitempty"`
	DeletedAt 		time.Time 		`json:"deleted_at,omitempty"`
}

type TicketReq struct {
	TrackId       	string    	`json:"track_id,omitempty"`
	AirplaneId     	string    	`json:"airplane_id,omitempty"`
	Datetime   		string 		`json:"datetime,omitempty"`
	Price    		int 		`json:"price,omitempty"`
}

type TicketRes struct {
	Id            		string   		`json:"id"`
	Track       		*TrackRes    	`json:"track,omitempty"`
	Airplane     		*AirplaneRes    	`json:"airplane,omitempty"`
	ArrivalDatetime   	time.Time 		`json:"arrival_datetime,omitempty"`
	DepartureDatetime   time.Time 		`json:"departure_datetime,omitempty"`
	Price    			int 			`json:"price,omitempty"`
	CreatedAt 			time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 			time.Time 		`json:"updated_at,omitempty"`
	DeleteAt 			bool 			`json:"deleted_at,omitempty"`
}

type TicketRepositoryPostgres interface {
	Insert(ticket *Ticket)(string, error)
	CheckTicketExist(trackId uuid.UUID, airplaneId uuid.UUID, datetime time.Time) error
	Delete(id string) error
	VerifyTicketAvailable(idTicket string) error
	Update(idTicket string, ticket *Ticket) error
	GetById(idTicket string)(*Ticket, *Track, *Airplane, error)
	List()([]*Ticket, error)
}
type TicketRepositoryElasticsearch interface {
	Insert(idTicket string, ticket *TicketES) error
	Update(idTicket string, ticket *TicketES) error
}

type TicketUseCase interface {
	Add(payload *TicketReq)(string, error)
	Delete(id string) error
	Update(idTicket string, ticket *TicketReq) error
	GetById(idTicket string)(*TicketRes, error)
	List()([]*TicketRes, error)
}