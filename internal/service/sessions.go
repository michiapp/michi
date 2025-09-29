package service

import (
	"context"

	"github.com/OrbitalJin/michi/internal/cache"
	"github.com/OrbitalJin/michi/internal/sqlc"
)

type SessionService struct {
	q     *sqlc.Queries
	cache *cache.Cache[string, sqlc.Session]
}

func NewSessionService(q *sqlc.Queries) *SessionService {
	return &SessionService{
		q:     q,
		cache: cache.New[string, sqlc.Session](),
	}
}

func (service *SessionService) Insert(ctx context.Context, alias string, urls []string) error {

	session, err := service.q.InsertSession(ctx, alias)

	if err != nil {
		return err
	}

	service.cache.Store(session.Alias, session)

	return nil
}

func (service *SessionService) GetFromAlias(ctx context.Context, alias string) (sqlc.Session, error) {
	shortcut, ok := service.cache.Load(alias)

	if ok {
		return shortcut, nil
	}

	shortcut, err := service.q.GetSessionByAlias(ctx, alias)

	if err != nil {
		return sqlc.Session{}, err
	}

	service.cache.Store(alias, shortcut)

	return shortcut, nil
}

func (service *SessionService) GetAll(ctx context.Context) ([]sqlc.Session, error) {
	return service.q.ListSessions(ctx)
}

func (service *SessionService) Delete(ctx context.Context, id int64) error {
	return service.q.DeleteSession(ctx, id)
}

func (service *SessionService) DeleteFromAlias(ctx context.Context, alias string) error {
	return service.q.DeleteSessionByAlias(ctx, alias)
}
