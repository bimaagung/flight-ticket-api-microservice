package domain

import (
	"time"

	"github.com/google/uuid"
)

type Airplane struct {
	Id            		uuid.UUID   		`json:"id"`
	FlightCode       	string    		`json:"flight_code,omitempty"`
	Seats    			int    			`json:"seats,omitempty"`
	Type    			string 			`json:"type,omitempty"`
	ProductionDate 		time.Time 		`json:"production_date,omitempty"`
	Factory    			string    		`json:"factory,omitempty"`
	CreatedAt 			time.Time 		`json:"created_at,omitempty"`
	UpdatedAt 			time.Time 		`json:"updated_at,omitempty"`
	DeleteAt 			time.Time 		`json:"deleted_at,omitempty"`
}

type AirplaneRes struct {
	Id            		string   	`json:"id"`
	FlightCode       	string    	`json:"flight_code,omitempty"`
	Seats    			int    		`json:"seats,omitempty"`
}
