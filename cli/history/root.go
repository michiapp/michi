package history

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(ctx context.Context, service *service.HistoryService) *v2.Command {
	return &v2.Command{
		Name:    "history",
		Usage:   "to manage history",
		Aliases: []string{"hs"},
		Subcommands: []*v2.Command{
			list(ctx, service),
			delete(ctx, service),
		},
	}
}

func fzf(history []sqlc.History, prompt string) *sqlc.History {
	index, err := fuzzy.FindMulti(
		history,
		func(i int) string {
			return history[i].Query

		},
		fuzzy.WithHeader("History - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Query: %s \nProvider: (%s) \nTimeStamp: %s",
				history[i].Query,
				history[i].ProviderTag,
				history[i].Timestamp,
			)
		}))

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &history[index[0]]
	}

	return nil
}
