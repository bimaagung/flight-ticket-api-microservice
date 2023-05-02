package mockspostgres

import (
	"ticket-service/domain"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type TicketPostgresRepository struct {
	mock.Mock
}

func (m *TicketPostgresRepository) Insert(ticket *domain.Ticket)(string, error){
	ret := m.Called(ticket)

	var r0 string

	if rf, ok := ret.Get(0).(func(*domain.Ticket) string); ok {
		r0 = rf(ticket)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(*domain.Ticket) error); ok {
		r1 = rf(ticket)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *TicketPostgresRepository) CheckTicketExist(trackId uuid.UUID, airplaneId uuid.UUID, datetime time.Time)error{
	ret := m.Called(trackId, airplaneId, datetime)

	var r0 error

	if rf, ok := ret.Get(0).(func(trackId uuid.UUID, airplaneId uuid.UUID, datetime time.Time) error); ok {
		r0 = rf(trackId, airplaneId, datetime)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketPostgresRepository) Delete(idTicket string)error{
	ret := m.Called(idTicket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string) error); ok {
		r0 = rf(idTicket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketPostgresRepository) VerifyTicketAvailable(idTicket string)error{
	ret := m.Called(idTicket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string) error); ok {
		r0 = rf(idTicket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketPostgresRepository) Update(idTicket string, ticket *domain.Ticket)error{
	ret := m.Called(idTicket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string) error); ok {
		r0 = rf(idTicket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketPostgresRepository) GetById(idTicket string)(*domain.Ticket, *domain.Track, *domain.Airplane, error){
	ret := m.Called(idTicket)

	var r0 *domain.Ticket

	if rf, ok := ret.Get(0).(func(idTicket string) *domain.Ticket); ok {
		r0 = rf(idTicket)
	}else {
		r0 = ret.Get(0).(*domain.Ticket)
	}

	var r1 *domain.Track

	if rf, ok := ret.Get(1).(func(idTicket string) *domain.Track); ok {
		r1 = rf(idTicket)
	}else {
		r1 = ret.Get(1).(*domain.Track)
	}

	var r2 *domain.Airplane

	if rf, ok := ret.Get(2).(func(idTicket string) *domain.Airplane); ok {
		r2 = rf(idTicket)
	}else {
		r2 = ret.Get(2).(*domain.Airplane)
	}

	var r3 error

	if rf, ok := ret.Get(3).(func(idTicket string) error); ok {
		r3 = rf(idTicket)
	}else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

func (m *TicketPostgresRepository) List()([]*domain.Ticket, error){
	ret := m.Called()

	var r0 []*domain.Ticket

	if rf, ok := ret.Get(0).(func()[]*domain.Ticket); ok {
		r0 = rf()
	}else {
		r0 = ret.Get(0).([]*domain.Ticket)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}