package mockses

import (
	"ticket-service/domain"

	"github.com/stretchr/testify/mock"
)

type TicketESRepository struct {
	mock.Mock
}

func (m *TicketESRepository) Insert(idTicket string, ticket *domain.TicketRes)error{
	ret := m.Called(idTicket, ticket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string, ticket *domain.TicketRes) error); ok {
		r0 = rf(idTicket, ticket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketESRepository) Update(idTicket string, ticket *domain.TicketRes)error{
	ret := m.Called(idTicket, ticket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string, ticket *domain.TicketRes) error); ok {
		r0 = rf(idTicket, ticket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketESRepository) Delete(idTicket string)error{
	ret := m.Called(idTicket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string) error); ok {
		r0 = rf(idTicket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TicketESRepository) Search(payloadSearch string)([]*domain.TicketRes, error){
	ret := m.Called(payloadSearch)

	var r0 []*domain.TicketRes

	if rf, ok := ret.Get(0).(func(payloadSearch string)[]*domain.TicketRes); ok {
		r0 = rf(payloadSearch)
	}else {
		r0 = ret.Get(0).([]*domain.TicketRes)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}