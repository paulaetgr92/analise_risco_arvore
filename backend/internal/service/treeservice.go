package service

import (
	db "arvore/db/sqlc"
	"arvore/internal/model"
	"arvore/internal/repository"
	"context"
	"database/sql"
	"errors"
	"time"
)

type TreeService struct {
	treeRepo repository.TreeRepositoryInterface
	riskRepo repository.RiskAssessmentRepositoryInterface
}

func NewTreeService(
	treeRepo repository.TreeRepositoryInterface,
	riskRepo repository.RiskAssessmentRepositoryInterface,
) *TreeService {
	return &TreeService{
		treeRepo: treeRepo,
		riskRepo: riskRepo,
	}
}

func (s *TreeService) CreateTreeService(ctx context.Context, arg model.Tree) (model.TreeResponse, error) {

	if arg.Latitude == 0 || arg.Longitude == 0 {
		return model.TreeResponse{}, errors.New("latitude e longitude são obrigatórias")
	}

	if arg.Height <= 0 {
		return model.TreeResponse{}, errors.New("altura inválida")
	}

	dbParams := db.CreateTreeParams{
		Latitude:  arg.Latitude,
		Longitude: arg.Longitude,
		Species: sql.NullString{
			String: arg.Species,
			Valid:  arg.Species != "",
		},
		Height: sql.NullFloat64{
			Float64: arg.Height,
			Valid:   true,
		},
	}

	// 🔥 PEGAR O RETORNO DO BANCO
	treeDB, err := s.treeRepo.CreateTreeRepository(ctx, dbParams)
	if err != nil {
		return model.TreeResponse{}, err
	}

	// 🔥 MAPEAR PARA RESPONSE
	response := model.TreeResponse{
		Id:        treeDB.ID,
		Latitude:  treeDB.Latitude,
		Longitude: treeDB.Longitude,
		Species:   treeDB.Species.String,
		Height:    treeDB.Height.Float64,
	}

	return response, nil
}
func (s *TreeService) GetTreeByIdService(ctx context.Context, id int64) (model.TreeResponse, error) {
	tree, err := s.treeRepo.GetTreeByIDRepository(ctx, id)
	if err != nil {
		return model.TreeResponse{}, err
	}
	response := model.TreeResponse{
		Id:        tree.ID,
		Latitude:  tree.Latitude,
		Longitude: tree.Longitude,
		Species:   tree.Species.String,
		Height:    tree.Height.Float64,
		Diameter:  tree.Height.Float64,
		Age:       int(tree.Age.Int32),
		Health:    tree.Health.String,
		CreatedAt: time.Time{},
	}
	return response, nil
}

func (s *TreeService) LisTreesByBoundingBoxService(ctx context.Context, arg model.ListTreesByBoundingBoxRequest) ([]model.ListTreesByBoundingBoxResponse, error) {

	params := db.ListTreesByBoundingBoxParams{
		Latitude:    arg.Latitude,
		Latitude_2:  arg.Latitude_2,
		Longitude:   arg.Longitude,
		Longitude_2: arg.Longitude_2,
	}
	_, err := s.treeRepo.LisTreeByBoundingBoxRepository(ctx, params)
	if err != nil {
		return []model.ListTreesByBoundingBoxResponse{}, err
	}
	return []model.ListTreesByBoundingBoxResponse{}, nil
}

func (s *TreeService) UpdateTreeService(ctx context.Context, arg model.UpdateTreeRequest) (model.TreeResponse, error) {
	if arg.ID == 0 {
		return model.TreeResponse{}, errors.New("id is required")
	}
	request := db.UpdateTreeParams{
		Latitude:  arg.Latitude,
		Longitude: arg.Longitude,
		Species: sql.NullString{
			String: arg.Species,
			Valid:  true,
		},
		ID: arg.ID,
	}
	_, err := s.treeRepo.UpdateTreeRepository(ctx, request)
	if err != nil {
		return model.TreeResponse{}, err
	}
	return model.TreeResponse{}, nil
}

func (s *TreeService) DeleteTreeService(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is required")
	}
	err := s.treeRepo.DeleteTreeRepository(ctx, id)
	if err != nil {
		return err
	}
	return err

}

func (s *TreeService) ListTreesService(ctx context.Context) ([]model.TreeResponse, error) {

	result, err := s.treeRepo.ListTreesRepository(ctx)
	if err != nil {
		return nil, err
	}

	var response []model.TreeResponse

	for _, tree := range result {
		response = append(response, model.TreeResponse{
			Id:        tree.ID,
			Latitude:  tree.Latitude,
			Longitude: tree.Longitude,
			Species:   tree.Species.String,
			Height:    tree.Height.Float64,
			Diameter:  tree.Diameter.Float64,
			Age:       int(tree.Age.Int32),
			Health:    tree.Health.String,
			CreatedAt: tree.CreatedAt.Time,
		})
	}

	return response, nil
}

func (s *TreeService) ListPotencialRiskTree(ctx context.Context) (model.TreeResponse, error) {
	err, _ := s.treeRepo.ListPotentialRiskTreesRepository(ctx)
	if err != nil {
		return model.TreeResponse{}, nil
	}
	return model.TreeResponse{}, nil
}
