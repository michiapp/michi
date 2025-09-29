package bangs

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	"github.com/urfave/cli/v2"
)

func Root(ctx context.Context, service *service.SPService) *cli.Command {
	return &cli.Command{
		Name:  "bangs",
		Usage: "to manage bangs",
		Subcommands: []*cli.Command{
			list(ctx, service),
			delete(ctx, service),
			create(ctx, service),
		},
	}
}

func fzf(bangs []sqlc.SearchProvider, prompt string) *sqlc.SearchProvider {
	index, err := fuzzy.FindMulti(
		bangs,
		func(i int) string {
			return bangs[i].Domain
		},
		fuzzy.WithHeader("Bangs - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			provider := bangs[i]

			return fmt.Sprintf(
				"Site Name: %s \n"+
					"Tag: %s \n"+
					"Category: %s \n"+
					"Subcategory: %s \n"+
					"Domain: %s \n"+
					"Rank: %d \n",
				provider.SiteName,
				provider.Tag,
				provider.Category,
				provider.Subcategory,
				provider.Domain,
				provider.Rank,
			)
		}),
	)

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &bangs[index[0]]
	}

	return nil
}
