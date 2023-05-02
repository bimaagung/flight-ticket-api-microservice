package mockspostgresrepository

import (
	"ticket-service/domain"

	"github.com/stretchr/testify/mock"
)

type AirplanePostgresRepository struct {
	mock.Mock
}

func (m *AirplanePostgresRepository) VerifyAirplaneAvailable(idAirplane string)error{
	ret := m.Called(idAirplane)

	var r0 error

	if rf, ok := ret.Get(0).(func(idAirplane string) error); ok {
		r0 = rf(idAirplane)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *AirplanePostgresRepository) Insert(airplane *domain.Airplane)(string, error){
	ret := m.Called(airplane)

	var r0 string

	if rf, ok := ret.Get(0).(func(*domain.Airplane) string); ok {
		r0 = rf(airplane)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(*domain.Airplane) error); ok {
		r1 = rf(airplane)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *AirplanePostgresRepository) CheckAirplaneExist(flightCode string)error{
	ret := m.Called(flightCode)

	var r0 error

	if rf, ok := ret.Get(0).(func(flightCode string) error); ok {
		r0 = rf(flightCode)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *AirplanePostgresRepository) List()([]*domain.Airplane, error){
	ret := m.Called()

	var r0 []*domain.Airplane

	if rf, ok := ret.Get(0).(func()[]*domain.Airplane); ok {
		r0 = rf()
	}else {
		r0 = ret.Get(0).([]*domain.Airplane)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}