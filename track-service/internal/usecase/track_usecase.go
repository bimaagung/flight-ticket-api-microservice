package usecase

import "track-service/domain"

func NewTrackUseCase(trackRepositoryPostgres domain.TrackRepositoryPostgres) domain.TrackUseCase {
	return &trackUseCaseImpl{
		TrackRepositoryPostgres: trackRepositoryPostgres,
	}
}

type trackUseCaseImpl struct {
	TrackRepositoryPostgres domain.TrackRepositoryPostgres
}

func(useCase *trackUseCaseImpl) Add(payload *domain.TrackReq)(string, error){
	
	track := &domain.Track{
		Arrival: payload.Arrival,
		Departure: payload.Departure,
		LongFlight: payload.LongFlight,
	}

	err := useCase.TrackRepositoryPostgres.CheckTrackExist(payload.Arrival, payload.Departure)

	if err != nil {
		return "", err
	}

	id, err := useCase.TrackRepositoryPostgres.Insert(track)

	if err != nil {
		return "", err
	}

	return id, nil
}

func(useCase *trackUseCaseImpl) GetById(id string)(*domain.TrackRes, error){

	track, err := useCase.TrackRepositoryPostgres.GetById(id)

	if err != nil {
		return nil, err
	}

	result := &domain.TrackRes{
		Id: track.Id.String(),
		Arrival: track.Arrival,
		Departure: track.Departure,
		LongFlight: track.LongFlight,
		CreatedAt: track.CreatedAt,
		UpdatedAt: track.UpdatedAt,
	}

	return result, nil
}

func(useCase *trackUseCaseImpl) List()([]*domain.TrackRes, error){

	track, err := useCase.TrackRepositoryPostgres.List()

	if err != nil {
		return nil, err
	}

	var tracks []*domain.TrackRes

	for _, v := range track {
		var track = domain.TrackRes{
			Id: v.Id.String(),
			Arrival: v.Arrival,
			Departure: v.Departure,
			LongFlight: v.LongFlight,
			CreatedAt: v.CreatedAt,
		}

		tracks = append(tracks, &track)
		
	}


	return tracks, nil
}

func(useCase *trackUseCaseImpl) Delete(id string) error {

	err := useCase.TrackRepositoryPostgres.VerifyTrackAvailable(id)
	if err != nil {
		return err
	}
	
	err = useCase.TrackRepositoryPostgres.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func(useCase *trackUseCaseImpl) Update(idTrack string, payload *domain.TrackReq)(*domain.TrackRes, error){
	
	track := &domain.Track{
		Arrival: payload.Arrival,
		Departure: payload.Departure,
		LongFlight: payload.LongFlight,
	}

	err := useCase.TrackRepositoryPostgres.VerifyTrackAvailable(idTrack)

	if err != nil {
		return nil, err
	}

	trackRes, err := useCase.TrackRepositoryPostgres.Update(idTrack, track)

	if err != nil {
		return nil, err
	}

	result := &domain.TrackRes{
		Id: trackRes.Id.String(),
		Arrival: trackRes.Arrival,
		Departure: trackRes.Departure,
		LongFlight: trackRes.LongFlight,
		CreatedAt: trackRes.CreatedAt,
		UpdatedAt: trackRes.UpdatedAt,
	}

	return result, nil
}