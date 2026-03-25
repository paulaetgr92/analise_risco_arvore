package repository

import (
	db "arvore/db/sqlc"
	"context"
	"database/sql"
)

type RiskAssessmentRepositoryInterface interface {
	CreateNewRiskAssesmentRepository(ctx context.Context, arg db.CreateRiskAssessmentParams) (db.RiskAssessment, error)
	GetRiskByIdRepository(ctx context.Context, id int64) (db.RiskAssessment, error)
	ListRiskByTreeRepository(ctx context.Context, treeId int32) ([]db.RiskAssessment, error)
	GetLatesteTreeByRiskAssessmentRepository(ctx context.Context, treeId int32) (db.RiskAssessment, error)
	ListHighRiskTreesRepository(ctx context.Context, treeId int32) ([]db.ListHighRiskTreesRow, error)
	ListMediumOrHighRiskTreesRepository(ctx context.Context, treeId int32) ([]db.ListMediumOrHighRiskTreesRow, error)
}

type TreeRepositoryInterface interface {
	CreateTreeRepository(ctx context.Context, arg db.CreateTreeParams) (db.Tree, error)
	GetTreeByIDRepository(ctx context.Context, id int64) (db.Tree, error)
	ListTreesRepository(ctx context.Context) ([]db.Tree, error)
	LisTreeByBoundingBoxRepository(ctx context.Context, arg db.ListTreesByBoundingBoxParams) ([]db.Tree, error)
	UpdateTreeRepository(ctx context.Context, arg db.UpdateTreeParams) (db.Tree, error)
	DeleteTreeRepository(ctx context.Context, id int64) error
	ListPotentialRiskTreesRepository(
		ctx context.Context,
	) ([]db.Tree, error)
}

type WeatherRepositoryInterface interface {
	CreateWeatherEventsRepository(ctx context.Context, arg db.CreateWeatherEventParams) (db.WeatherEvent, error)
	GetNewWeatherEventByIdRepository(ctx context.Context, id int64) (db.WeatherEvent, error)
	ListWeatherEventsRepository(ctx context.Context) ([]db.WeatherEvent, error)
	ListWeatherEventsByRegionRepository(ctx context.Context, region string) ([]db.WeatherEvent, error)
	ListRecentWeatherEventsRepository(ctx context.Context) ([]db.WeatherEvent, error)
	ListSeveralWeatherEventsRepository(ctx context.Context) ([]db.WeatherEvent, error)
	ListHighWindEvenytesRepository(ctx context.Context, windSpeed float64) ([]db.WeatherEvent, error)
	ListHeavyRainEventsRepository(ctx context.Context, rainfallMm sql.NullFloat64) ([]db.WeatherEvent, error)
	GetLatestSevereWeatherByRegionRepository(ctx context.Context, region string) (db.WeatherEvent, error)
}
