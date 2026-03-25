package service

import (
	db "arvore/db/sqlc"
	"arvore/internal/model"
	"arvore/internal/repository"
	"context"
	"errors"
)

type WeatherEventService struct {
	weatherRepo repository.WeatherRepositoryInterface
}

func NewWeatherEventService(
	weatherRepo repository.WeatherRepositoryInterface,
) *WeatherEventService {
	return &WeatherEventService{
		weatherRepo: weatherRepo,
	}
}

func (s *WeatherEventService) CreateWeatherEvent(
	ctx context.Context,
	req model.WeatherEventRequest,
) (model.WeatherEventResponse, error) {

	if req.Region == "" {
		return model.WeatherEventResponse{}, errors.New("região é obrigatória")
	}

	if req.WindSpeed < 0 {
		return model.WeatherEventResponse{}, errors.New("velocidade do vento inválida")
	}

	dbParams := db.CreateWeatherEventParams{
		Region:    req.Region,
		WindSpeed: req.WindSpeed,
	}

	event, err := s.weatherRepo.CreateWeatherEventsRepository(ctx, dbParams)
	if err != nil {
		return model.WeatherEventResponse{}, err
	}

	return model.WeatherEventResponse{
		ID:        event.ID,
		Region:    event.Region,
		WindSpeed: event.WindSpeed,
	}, nil
}

func (s *WeatherEventService) ListWeatherEvents(
	ctx context.Context,
) ([]model.WeatherEventResponse, error) {

	events, err := s.weatherRepo.ListWeatherEventsRepository(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]model.WeatherEventResponse, 0, len(events))
	for _, e := range events {
		res = append(res, model.WeatherEventResponse{
			ID:        e.ID,
			Region:    e.Region,
			WindSpeed: e.WindSpeed,
		})
	}

	return res, nil
}

func (s *WeatherEventService) ListWeatherEventsByRegion(
	ctx context.Context,
	region string,
) ([]model.WeatherEventResponse, error) {

	if region == "" {
		return nil, errors.New("região inválida")
	}

	events, err := s.weatherRepo.ListWeatherEventsByRegionRepository(ctx, region)
	if err != nil {
		return nil, err
	}

	res := make([]model.WeatherEventResponse, 0, len(events))
	for _, e := range events {
		res = append(res, model.WeatherEventResponse{
			ID:        e.ID,
			Region:    e.Region,
			WindSpeed: e.WindSpeed,
		})
	}

	return res, nil
}

func (s *WeatherEventService) ListRecentWeatherEvents(
	ctx context.Context,
) ([]model.WeatherEventResponse, error) {

	events, err := s.weatherRepo.ListRecentWeatherEventsRepository(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]model.WeatherEventResponse, 0, len(events))
	for _, e := range events {
		res = append(res, model.WeatherEventResponse{
			ID:        e.ID,
			Region:    e.Region,
			WindSpeed: e.WindSpeed,
		})
	}

	return res, nil
}

func (s *WeatherEventService) GetLatestSevereWeatherByRegion(
	ctx context.Context,
	region string,
) (model.WeatherEventResponse, error) {

	if region == "" {
		return model.WeatherEventResponse{}, errors.New("região inválida")
	}

	event, err := s.weatherRepo.GetLatestSevereWeatherByRegionRepository(ctx, region)
	if err != nil {
		return model.WeatherEventResponse{}, err
	}

	return model.WeatherEventResponse{
		ID:        event.ID,
		Region:    event.Region,
		WindSpeed: event.WindSpeed,
	}, nil
}
