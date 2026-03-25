package repository

import (
	db "arvore/db/sqlc"
	"context"
)

type TreeNewRepository struct {
	*BaseRepository
}

func NewTreeRepository(base *BaseRepository) *TreeNewRepository {
	return &TreeNewRepository{BaseRepository: base}
}

func (r *TreeNewRepository) CreateTreeRepository(
	ctx context.Context,
	arg db.CreateTreeParams,
) (db.Tree, error) {

	if err := r.GetConnection(ctx); err != nil {
		return db.Tree{}, err
	}

	return r.Queries.CreateTree(ctx, arg)
}

func (r *TreeNewRepository) GetTreeByIDRepository(ctx context.Context, id int64) (db.Tree, error) {
	if err := r.GetConnection(ctx); err != nil {
		return db.Tree{}, err
	}
	return r.Queries.GetTreeByID(ctx, id)
}

func (r *TreeNewRepository) ListTreesRepository(ctx context.Context) ([]db.Tree, error) {
	if err := r.GetConnection(ctx); err != nil {
		return []db.Tree{}, err
	}
	return r.Queries.ListTrees(ctx)
}

func (r *TreeNewRepository) LisTreeByBoundingBoxRepository(ctx context.Context, arg db.ListTreesByBoundingBoxParams) ([]db.Tree, error) {
	if err := r.GetConnection(ctx); err != nil {
		return []db.Tree{}, err
	}
	return r.Queries.ListTreesByBoundingBox(ctx, arg)
}

func (r *TreeNewRepository) UpdateTreeRepository(ctx context.Context, arg db.UpdateTreeParams) (db.Tree, error) {
	if err := r.GetConnection(ctx); err != nil {
		return db.Tree{}, err
	}
	return r.Queries.UpdateTree(ctx, arg)
}

func (r *TreeNewRepository) DeleteTreeRepository(ctx context.Context, id int64) error {
	if err := r.GetConnection(ctx); err != nil {
		return err
	}
	return r.Queries.DeleteTree(ctx, id)
}

func (r *TreeNewRepository) ListPotentialRiskTreesRepository(
	ctx context.Context,
) ([]db.Tree, error) {

	if err := r.GetConnection(ctx); err != nil {
		return nil, err
	}

	return r.Queries.ListPotentialRiskTrees(ctx)
}
