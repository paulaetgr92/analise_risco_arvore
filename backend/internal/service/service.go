package service

import (
	"arvore/internal/model"
	"context"
)

type TreeServiceInterface interface {
	CreateTreeService(ctx context.Context, arg model.Tree) (model.TreeResponse, error)
	GetTreeByIdService(ctx context.Context, id int64) (model.TreeResponse, error)
	LisTreesByBoundingBoxService(ctx context.Context, arg model.ListTreesByBoundingBoxRequest) ([]model.ListTreesByBoundingBoxResponse, error)
	UpdateTreeService(ctx context.Context, arg model.UpdateTreeRequest) (model.TreeResponse, error)
	DeleteTreeService(ctx context.Context, id int64) error
	ListPotencialRiskTree(ctx context.Context) (model.TreeResponse, error)
}

type RiskAssesmentServiceInterface interface {
	CreateRiskAssessment(
		ctx context.Context,
		treeID int64,
		score int32,
		level string,
	) (model.RiskAssessmentResponse, error)
	GetRiskByID(
		ctx context.Context,
		id int64,
	) (model.RiskAssessmentResponse, error)
	ListRiskByTree(
		ctx context.Context,
		treeID int64,
	) ([]model.RiskAssessmentResponse, error)
	GetLatestRiskByTree(
		ctx context.Context,
		treeID int64,
	) (model.RiskAssessmentResponse, error)
	ListHighRiskTrees(
		ctx context.Context,
	) ([]model.HighRiskTreeResponse, error)
	ListMediumOrHighRiskTrees(
		ctx context.Context,
	) ([]model.MediumOrHighRiskTreeResponse, error)
}

type WeatherEventsServiceInterface interface {
	CreateWeatherEvent(
		ctx context.Context,
		req model.WeatherEventRequest,
	) (model.WeatherEventResponse, error)
	ListWeatherEvents(
		ctx context.Context,
	) ([]model.WeatherEventResponse, error)
	ListWeatherEventsByRegion(
		ctx context.Context,
		region string,
	) ([]model.WeatherEventResponse, error)
	ListRecentWeatherEvents(
		ctx context.Context,
	) ([]model.WeatherEventResponse, error)
	GetLatestSevereWeatherByRegion(
		ctx context.Context,
		region string,
	) (model.WeatherEventResponse, error)
}
