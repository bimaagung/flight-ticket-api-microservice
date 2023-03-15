package domain

import (
	"time"
)

type Airplane struct {
	Id            		string   	`json:"id"`
	FlightCode       	string    	`json:"flight_code,omitempty"`
	Seats    			int    		`json:"seats,omitempty"`
	Type    			string 		`json:"type,omitempty"`
	Production 			time.Time 	`json:"production,omitempty"`
	Factory    			string    	`json:"factory,omitempty"`
	CreatedAt 			time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 			time.Time 	`json:"updated_at,omitempty"`
	DeleteAt 			time.Time 	`json:"deleted_at,omitempty"`
}

type AirplaneReq struct {
	FlightCode       	string    	`json:"flight_code,omitempty"`
	Seats    			int    	`json:"seats,omitempty"`
	Type    			string 		`json:"type,omitempty"`
	Production 			time.Time 	`json:"production,omitempty"`
	Factory    			string    	`json:"factory,omitempty"`
}

type AirplaneRes struct {
	Id            		string   	`json:"id"`
	FlightCode       	string    	`json:"flight_code,omitempty"`
	Seats    			int    	`json:"seats,omitempty"`
	Type    			string 		`json:"type,omitempty"`
	Production 			time.Time 	`json:"production,omitempty"`
	Factory    			string    	`json:"factory,omitempty"`
	CreatedAt 			time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 			time.Time 	`json:"updated_at,omitempty"`
}

type AirplaneRepositoryMysql interface {
	Insert(track *Airplane)(string, error)
	CheckAirplaneExist(flightCode string) error
	GetById(id string)(*Airplane, error)
	List()([]*Airplane, error)
	Delete(id string) error
	VerifyAirplaneAvailable(idAirplane string) error
	Update(idAirplane string, track *Airplane)(*Airplane, error)
}

type AirplaneUseCase interface {
	Add(payload *AirplaneReq)(string, error)
	GetById(idAirplane string)(*AirplaneRes, error)
	List()([]*AirplaneRes, error)
	Delete(id string) error
	Update(idAirplane string, track *AirplaneReq)(*AirplaneRes, error)
}