package repository

import (
	db "arvore/db/sqlc"
	"context"
)

type RiskAssesmentNewRepository struct {
	*BaseRepository
}

func NewRiskAssesmentRepository(base *BaseRepository) *RiskAssesmentNewRepository {
	return &RiskAssesmentNewRepository{BaseRepository: base}
}

func (r *RiskAssesmentNewRepository) CreateNewRiskAssesmentRepository (ctx context.Context, arg db.CreateRiskAssessmentParams) (db.RiskAssessment, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.RiskAssessment{}, err
	}
	return r.Queries.CreateRiskAssessment(ctx, arg)
}

func (r* RiskAssesmentNewRepository) GetRiskByIdRepository (ctx context.Context, id int64) (db.RiskAssessment, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.RiskAssessment{}, err
	}
	return r.Queries.GetRiskByID(ctx, id)
}

func (r*RiskAssesmentNewRepository)ListRiskByTreeRepository (ctx context.Context, treeId int32) ([]db.RiskAssessment, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return []db.RiskAssessment{}, err
	}
	return r.Queries.ListRiskByTree(ctx, treeId)
}
func (r*RiskAssesmentNewRepository) GetLatesteTreeByRiskAssessmentRepository (ctx context.Context, treeId int32) (db.RiskAssessment, error) {
	err:= r.GetConnection(ctx)
	if err != nil {
		return db.RiskAssessment{}, err

	}
	return r.Queries.GetLatestRiskByTree(ctx, treeId)
}

func (r * RiskAssesmentNewRepository) ListHighRiskTreesRepository (ctx context.Context, treeId int32) ([]db.ListHighRiskTreesRow, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return []db.ListHighRiskTreesRow{}, err
	}
	return r.Queries.ListHighRiskTrees(ctx, treeId)
}

func ( r * RiskAssesmentNewRepository) ListMediumOrHighRiskTreesRepository (ctx context.Context, treeId int32) ([]db.ListMediumOrHighRiskTreesRow, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return []db.ListMediumOrHighRiskTreesRow{}, err
	}
	return r.Queries.ListMediumOrHighRiskTrees(ctx, treeId)
}