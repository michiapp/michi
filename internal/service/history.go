package service

import (
	"context"

	"github.com/OrbitalJin/michi/internal/sqlc"
)

type HistoryService struct {
	q *sqlc.Queries
}

func NewHistoryService(q *sqlc.Queries) *HistoryService {
	return &HistoryService{
		q: q,
	}
}

func (service *HistoryService) Insert(ctx context.Context, params sqlc.InsertHistoryEntryParams) error {
	return service.q.InsertHistoryEntry(ctx, params)
}

func (service *HistoryService) GetRecentHistory(ctx context.Context, limit int64) ([]sqlc.History, error) {
	return service.q.GetRecentHistory(ctx, limit)
}

func (service *HistoryService) GetAllHistory(ctx context.Context) ([]sqlc.History, error) {
	return service.q.ListHistory(ctx)
}

func (service *HistoryService) DeleteEntry(ctx context.Context, id int64) error {
	return service.q.DeleteHistoryEntry(ctx, id)
}
