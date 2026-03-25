package repository

import (
	db "arvore/db/sqlc"
	"context"
	"database/sql"
)

type WeatherNewRepository struct {
*BaseRepository
}

func NewWeatherRepository(base *BaseRepository) *WeatherNewRepository {
	return &WeatherNewRepository{BaseRepository: base}
}

func (r* WeatherNewRepository) CreateWeatherEventsRepository (ctx context.Context, arg db.CreateWeatherEventParams) (db.WeatherEvent, error) {
	if err := r.GetConnection(ctx);
	err != nil {
		return db.WeatherEvent{}, err

	}
	return r.Queries.CreateWeatherEvent(ctx,arg)
}

func (r *WeatherNewRepository) GetNewWeatherEventByIdRepository ( ctx context.Context, id int64 ) (db.WeatherEvent, error) {
	if err := r.GetConnection(ctx);
	err != nil {
		return db.WeatherEvent{}, err
	}
	return r.Queries.GetWeatherEventByID( ctx, id)
}

func (r *WeatherNewRepository) ListWeatherEventsRepository (ctx context.Context) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx);
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListWeatherEvents(ctx)
}

func (r *WeatherNewRepository) ListWeatherEventsByRegionRepository (ctx context.Context, region string) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx);
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListWeatherEventsByRegion(ctx, region)
}
func (r *WeatherNewRepository) ListRecentWeatherEventsRepository (ctx context.Context) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx);
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListRecentWeatherEvents(ctx)
}

func (r *WeatherNewRepository) ListSeveralWeatherEventsRepository (ctx context.Context) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx)
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListSevereWeatherEvents(ctx)
}

func (r* WeatherNewRepository) ListHighWindEvenytesRepository (ctx context.Context, windSpeed float64) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx)
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListHighWindEvents(ctx, windSpeed)
}

func (r* WeatherNewRepository) ListHeavyRainEventsRepository (ctx context.Context, rainfallMm sql.NullFloat64) ([]db.WeatherEvent, error) {
	if err := r.GetConnection(ctx)
	err != nil {
		return []db.WeatherEvent{}, err
	}
	return r.Queries.ListHeavyRainEvents(ctx, rainfallMm)
}

func (r *WeatherNewRepository)GetLatestSevereWeatherByRegionRepository (ctx context.Context, region string) (db.WeatherEvent, error) {
	if err := r.GetConnection(ctx)
	err != nil {
		return db.WeatherEvent{}, err
	}
	return r.Queries.GetLatestSevereWeatherByRegion(ctx, region)
}