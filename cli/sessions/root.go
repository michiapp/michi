package sessions

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(ctx context.Context, service *service.SessionService) *v2.Command {
	return &v2.Command{
		Name:    "sessions",
		Usage:   "Manage sessions",
		Aliases: []string{"sesh"},
		Subcommands: []*v2.Command{
			list(ctx, service),
			add(ctx, service),
			delete(ctx, service),
		},
	}
}

func fzf(sessions []models.SessionWithUrls, prompt string) *sqlc.Session {
	index, err := fuzzy.FindMulti(
		sessions,
		func(i int) string {
			return sessions[i].Session.Alias

		},
		fuzzy.WithHeader("Sessions - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			urlsStr := ""
			for _, url := range sessions[i].Urls {
				urlsStr += url.Url + "\n"
			}
			return fmt.Sprintf("Alias: %s \nURLs:\n%s \nCreated At: %s",
				sessions[i].Session.Alias,
				urlsStr,
				sessions[i].Session.CreatedAt,
			)
		}))

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &sessions[index[0]].Session
	}

	return nil
}
