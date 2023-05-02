package mockspostgres

import (
	"ticket-service/domain"

	"github.com/stretchr/testify/mock"
)

type TrackPostgresRepository struct {
	mock.Mock
}

func (m *TrackPostgresRepository) CheckTrackExist(arrival string, depature string)error{
	ret := m.Called(arrival, depature)

	var r0 error

	if rf, ok := ret.Get(0).(func(arrival string, depature string) error); ok {
		r0 = rf(arrival, depature)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *TrackPostgresRepository) VerifyTrackAvailable(idTrack string)error{
	ret := m.Called(idTrack)

	var r0 error

	if rf, ok := ret.Get(0).(func(idTrack string) error); ok {
		r0 = rf(idTrack)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}


func (m *TrackPostgresRepository) Insert(track *domain.Track)(string, error){
	ret := m.Called(track)

	var r0 string

	if rf, ok := ret.Get(0).(func(*domain.Track) string); ok {
		r0 = rf(track)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(*domain.Track) error); ok {
		r1 = rf(track)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}


func (m *TrackPostgresRepository) List()([]*domain.Track, error){
	ret := m.Called()

	var r0 []*domain.Track

	if rf, ok := ret.Get(0).(func()[]*domain.Track); ok {
		r0 = rf()
	}else {
		r0 = ret.Get(0).([]*domain.Track)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

