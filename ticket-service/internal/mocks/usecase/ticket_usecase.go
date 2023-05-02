package mockusecase

import (
	"ticket-service/domain"

	"github.com/stretchr/testify/mock"
)

type TicketUseCase struct {
	mock.Mock
}

func(m *TicketUseCase) Add(payload *domain.TicketReq)(string, error) {
	ret := m.Called(payload)

	var r0 string

	if rf, ok := ret.Get(0).(func(*domain.TicketReq) string); ok {
		r0 = rf(payload)
	}else{
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(*domain.TicketReq) error); ok {
		r1 = rf(payload)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(m *TicketUseCase) Delete(id string)error {
	ret := m.Called(id)

	var r0 error

	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func(m *TicketUseCase) Update(idTicket string, ticket *domain.TicketReq) error {
	ret := m.Called(idTicket, ticket)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTicket string, ticket *domain.TicketReq) error); ok {
		r0 = rf(idTicket, ticket)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func(m *TicketUseCase) GetById(idTicket string)(*domain.TicketRes, error) {
	ret := m.Called(idTicket)

	var r0 *domain.TicketRes

	if rf, ok := ret.Get(0).(func(string) *domain.TicketRes); ok {
		r0 = rf(idTicket)
	}else{
		r0 = ret.Get(0).(*domain.TicketRes)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idTicket)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(m *TicketUseCase) List()([]*domain.TicketRes, error) {
	ret := m.Called()

	var r0 []*domain.TicketRes

	if rf, ok := ret.Get(0).(func() []*domain.TicketRes); ok {
		r0 = rf()
	}else{
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

func(m *TicketUseCase) Search(payloadSearch string)([]*domain.TicketRes, error) {
	ret := m.Called(payloadSearch)

	var r0 []*domain.TicketRes

	if rf, ok := ret.Get(0).(func(string) []*domain.TicketRes); ok {
		r0 = rf(payloadSearch)
	}else{
		r0 = ret.Get(0).([]*domain.TicketRes)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(payloadSearch)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}