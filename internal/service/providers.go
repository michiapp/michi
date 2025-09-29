package service

import (
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/OrbitalJin/michi/internal/cache"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/sqlc"
)

type SPService struct {
	q               *sqlc.Queries
	parser          *parser.Parser
	defaultProvider string
	cache           *cache.Cache[string, sqlc.SearchProvider]
}

func NewSearchProviderService(
	p *parser.Parser,
	q *sqlc.Queries,
	defaultProvider string,
) *SPService {

	return &SPService{
		parser:          p,
		q:               q,
		defaultProvider: defaultProvider,
		cache:           cache.New[string, sqlc.SearchProvider](),
	}
}

func (s *SPService) Insert(ctx context.Context, params sqlc.InsertProviderParams) error {
	params.Rank = 0
	return s.q.InsertProvider(ctx, params)
}

func (service *SPService) GetByTag(ctx context.Context, t string) (sqlc.SearchProvider, error) {
	provider, ok := service.cache.Load(t)

	if ok {
		return provider, nil
	}

	provider, err := service.q.GetProviderByTag(ctx, t)

	if err == nil {
		service.cache.Store(t, provider)
	}

	return provider, err
}

func (service *SPService) Collect(ctx context.Context, v string) ([]sqlc.SearchProvider, error) {
	result, err := service.parser.Collect(v)

	if err != nil {
		return nil, err
	}

	if len(result.Matches) == 0 {
		return nil, nil
	}

	var sps []sqlc.SearchProvider

	for _, tag := range result.Matches {
		p, err := service.GetByTag(ctx, tag)

		if err != nil {
			continue
		}

		sps = append(sps, p)
	}

	return sps, nil
}

func (service *SPService) CollectAndRank(ctx context.Context, v string) (
	*parser.Result,
	sqlc.SearchProvider,
	error,
) {
	result, err := service.parser.Collect(v)

	if err != nil {
		return nil, sqlc.SearchProvider{}, err
	}

	if len(result.Matches) == 0 {
		return result, sqlc.SearchProvider{}, nil
	}

	best := service.Rank(ctx, result)

	return result, best, nil
}

func (service *SPService) Rank(ctx context.Context, result *parser.Result) sqlc.SearchProvider {
	if result == nil {
		return sqlc.SearchProvider{}
	}

	var best sqlc.SearchProvider
	var bestRank int64 = -1

	for _, tag := range result.Matches {
		p, err := service.GetByTag(ctx, tag)

		if err != nil {
			continue
		}

		if p.Rank > bestRank {
			best = p
			bestRank = p.Rank
		}
	}

	return best
}

func (service *SPService) Resolve(
	query string,
	provider sqlc.SearchProvider,
) (sqlc.SearchProvider, *string, error) {

	encoded := url.QueryEscape(query)
	url := strings.Replace(provider.Url, "{{{s}}}", encoded, 1)
	return provider, &url, nil
}

func (service *SPService) ResolveWithFallback(ctx context.Context, query string) (sqlc.SearchProvider, *string, error) {
	p, err := service.GetByTag(ctx, service.defaultProvider)

	if err != nil {
		log.Println("no default provider available.")
		return sqlc.SearchProvider{}, nil, err
	}

	return service.Resolve(query, p)
}

func (service *SPService) GetAll(ctx context.Context) ([]sqlc.SearchProvider, error) {
	return service.q.ListProviders(ctx)
}

func (service *SPService) Delete(ctx context.Context, id int64) error {
	return service.q.DeleteProvider(ctx, id)
}
