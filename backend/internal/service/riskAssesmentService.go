package service

import (
	db "arvore/db/sqlc"
	"arvore/internal/model"
	"arvore/internal/repository"
	"context"
	"errors"
)

type RiskAssessmentService struct {
	riskRepo repository.RiskAssessmentRepositoryInterface
	treeRepo repository.TreeRepositoryInterface
}

func NewRiskAssessmentService(
	riskRepo repository.RiskAssessmentRepositoryInterface,
	treeRepo repository.TreeRepositoryInterface,
) *RiskAssessmentService {
	return &RiskAssessmentService{
		riskRepo: riskRepo,
		treeRepo: treeRepo,
	}
}

func (s *RiskAssessmentService) CreateRiskAssessment(
	ctx context.Context,
	treeID int64,
	score int32,
	level string,
) (model.RiskAssessmentResponse, error) {

	if treeID == 0 {
		return model.RiskAssessmentResponse{}, errors.New("tree_id inválido")
	}

	if score < 0 {
		return model.RiskAssessmentResponse{}, errors.New("score inválido")
	}

	if level == "" {
		return model.RiskAssessmentResponse{}, errors.New("nível inválido")
	}

	_, err := s.treeRepo.GetTreeByIDRepository(ctx, treeID)
	if err != nil {
		return model.RiskAssessmentResponse{}, errors.New("árvore não encontrada")
	}

	risk, err := s.riskRepo.CreateNewRiskAssesmentRepository(
		ctx,
		db.CreateRiskAssessmentParams{
			TreeID: int32(treeID),
			Score:  score,
			Level:  level,
		},
	)
	if err != nil {
		return model.RiskAssessmentResponse{}, err
	}

	return model.RiskAssessmentResponse{
		ID:           risk.ID,
		TreeID:       int64(risk.TreeID),
		Score:        risk.Score,
		Level:        risk.Level,
		CalculatedAt: risk.CalculatedAt,
	}, nil
}

func (s *RiskAssessmentService) GetRiskByID(
	ctx context.Context,
	id int64,
) (model.RiskAssessmentResponse, error) {

	risk, err := s.riskRepo.GetRiskByIdRepository(ctx, id)
	if err != nil {
		return model.RiskAssessmentResponse{}, err
	}

	return model.RiskAssessmentResponse{
		ID:           risk.ID,
		TreeID:       int64(risk.TreeID),
		Score:        risk.Score,
		Level:        risk.Level,
		CalculatedAt: risk.CalculatedAt,
	}, nil
}

func (s *RiskAssessmentService) ListRiskByTree(
	ctx context.Context,
	treeID int64,
) ([]model.RiskAssessmentResponse, error) {

	if treeID == 0 {
		return nil, errors.New("tree_id inválido")
	}

	risks, err := s.riskRepo.ListRiskByTreeRepository(ctx, int32(treeID))
	if err != nil {
		return nil, err
	}

	res := make([]model.RiskAssessmentResponse, 0, len(risks))
	for _, r := range risks {
		res = append(res, model.RiskAssessmentResponse{
			ID:           r.ID,
			TreeID:       int64(r.TreeID),
			Score:        r.Score,
			Level:        r.Level,
			CalculatedAt: r.CalculatedAt,
		})
	}

	return res, nil
}

func (s *RiskAssessmentService) GetLatestRiskByTree(
	ctx context.Context,
	treeID int64,
) (model.RiskAssessmentResponse, error) {

	if treeID == 0 {
		return model.RiskAssessmentResponse{}, errors.New("tree_id inválido")
	}

	risk, err := s.riskRepo.GetLatesteTreeByRiskAssessmentRepository(ctx, int32(treeID))
	if err != nil {
		return model.RiskAssessmentResponse{}, err
	}

	return model.RiskAssessmentResponse{
		ID:           risk.ID,
		TreeID:       int64(risk.TreeID),
		Score:        risk.Score,
		Level:        risk.Level,
		CalculatedAt: risk.CalculatedAt,
	}, nil
}

func (s *RiskAssessmentService) ListHighRiskTrees(
	ctx context.Context,
) ([]model.HighRiskTreeResponse, error) {

	rows, err := s.riskRepo.ListHighRiskTreesRepository(ctx, 0)
	if err != nil {
		return nil, err
	}

	res := make([]model.HighRiskTreeResponse, 0, len(rows))
	for _, r := range rows {

		tree := model.HighRiskTreeResponse{
			TreeID:    r.ID, // 👈 ID da árvore
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
			Score:     r.Score,
			Level:     r.Level,
		}

		if r.Species.Valid {
			tree.Species = r.Species.String
		}
		if r.Height.Valid {
			tree.Height = r.Height.Float64
		}
		if r.Diameter.Valid {
			tree.Diameter = r.Diameter.Float64
		}
		if r.Age.Valid {
			tree.Age = r.Age.Int32
		}
		if r.Health.Valid {
			tree.Health = r.Health.String
		}
		if r.Status.Valid {
			tree.Status = r.Status.String
		}
		if r.CreatedAt.Valid {
			tree.CreatedAt = r.CreatedAt.Time
		}

		res = append(res, tree)
	}

	return res, nil
}

func (s *RiskAssessmentService) ListMediumOrHighRiskTrees(
	ctx context.Context,
) ([]model.MediumOrHighRiskTreeResponse, error) {

	rows, err := s.riskRepo.ListMediumOrHighRiskTreesRepository(ctx, 0)
	if err != nil {
		return nil, err
	}

	res := make([]model.MediumOrHighRiskTreeResponse, 0, len(rows))
	for _, r := range rows {

		tree := model.MediumOrHighRiskTreeResponse{
			TreeID:    r.ID, // 👈 ID da árvore
			Latitude:  r.Latitude,
			Longitude: r.Longitude,
			Score:     r.Score,
			Level:     r.Level,
		}

		if r.Species.Valid {
			tree.Species = r.Species.String
		}
		if r.Height.Valid {
			tree.Height = r.Height.Float64
		}
		if r.Diameter.Valid {
			tree.Diameter = r.Diameter.Float64
		}
		if r.Age.Valid {
			tree.Age = r.Age.Int32
		}
		if r.Health.Valid {
			tree.Health = r.Health.String
		}
		if r.Status.Valid {
			tree.Status = r.Status.String
		}
		if r.CreatedAt.Valid {
			tree.CreatedAt = r.CreatedAt.Time
		}

		res = append(res, tree)
	}

	return res, nil
}
