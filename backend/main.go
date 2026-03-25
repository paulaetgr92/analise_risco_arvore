package main

import (
	"arvore/config"
	sqlc "arvore/db/sqlc"
	"arvore/internal/handler"
	"arvore/internal/repository"
	"arvore/internal/service"

	"github.com/labstack/echo/v4"
)
import "github.com/labstack/echo/v4/middleware"

func main() {

	e := echo.New()
	e.Use(middleware.CORS())

	dbConn := config.NewDatabase()

	// sqlc
	queries := sqlc.New(dbConn)

	// base repository
	baseRepo := repository.NewBaseRepository(queries)

	// repositories
	treeRepo := repository.NewTreeRepository(baseRepo)
	riskRepo := repository.NewRiskAssesmentRepository(baseRepo)

	// service
	treeService := service.NewTreeService(treeRepo, riskRepo)

	// handler
	treeHandler := handler.NewTreeHandler(treeService)

	// routes
	config.RegisterTreeRoutes(e, treeHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
