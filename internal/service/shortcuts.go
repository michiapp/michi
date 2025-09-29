package service

import (
	"context"

	"github.com/OrbitalJin/michi/internal/cache"
	"github.com/OrbitalJin/michi/internal/sqlc"
)

type ShortcutService struct {
	q     *sqlc.Queries
	cache *cache.Cache[string, sqlc.Shortcut]
}

func NewShortcutService(q *sqlc.Queries) *ShortcutService {
	return &ShortcutService{
		q:     q,
		cache: cache.New[string, sqlc.Shortcut](),
	}
}

func (service *ShortcutService) Insert(ctx context.Context, params sqlc.InsertShortcutParams) error {
	sc, err := service.q.InsertShortcut(ctx, params)

	if err != nil {
		return err
	}

	service.cache.Store(sc.Alias, sc)
	return nil
}

func (service *ShortcutService) GetFromAlias(ctx context.Context, alias string) (sqlc.Shortcut, error) {
	shortcut, ok := service.cache.Load(alias)

	if ok {
		return shortcut, nil
	}

	shortcut, err := service.q.GetShortcutByAlias(ctx, alias)

	if err != nil {
		return sqlc.Shortcut{}, err
	}

	service.cache.Store(alias, shortcut)

	return shortcut, nil
}

func (service *ShortcutService) GetAll(ctx context.Context) ([]sqlc.Shortcut, error) {
	return service.q.ListShortcuts(ctx)
}

func (service *ShortcutService) Delete(ctx context.Context, id int64) error {
	return service.q.DeleteShortcut(ctx, id)
}

func (service *ShortcutService) DeleteFromAlias(ctx context.Context, alias string) error {
	return service.q.DeleteShortcutFromAlias(ctx, alias)
}
