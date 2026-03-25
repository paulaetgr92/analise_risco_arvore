package handler

import "github.com/labstack/echo/v4"

type RiskAssesmentHandlerInterface interface {
	CreateRiskAssessment(c echo.Context) error
	GetLatestRiskByTree(c echo.Context) error
	ListRiskByTree(c echo.Context) error
}

type TreeHandlerInterface interface {
	CreateTreeHandler(c echo.Context) error
	ListTreesByBoundingBox(c echo.Context) error
	GetTreeByID(c echo.Context) error
	UpdateTree(c echo.Context) error
	DeleteTree(c echo.Context) error
	ListPotentialRiskTree(c echo.Context) error
}
