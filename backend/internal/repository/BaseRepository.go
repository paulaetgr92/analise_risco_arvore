package repository

import (
	db "arvore/db/sqlc"
	"context"
)

type BaseRepository struct {
	Queries *db.Queries
}

func NewBaseRepository(q *db.Queries) *BaseRepository {
	return &BaseRepository{
		Queries: q,
	}
}

func (b *BaseRepository) GetConnection(ctx context.Context) error {
	return nil
}
