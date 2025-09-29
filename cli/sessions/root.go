package sessions

import (
	"fmt"
	"strings"

	"github.com/OrbitalJin/michi/internal/service"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(service service.SessionServiceIface) *v2.Command {
	return &v2.Command{
		Name:    "sessions",
		Usage:   "Manage sessions",
		Aliases: []string{"sesh"},
		Subcommands: []*v2.Command{
			list(service),
			add(service),
			delete(service),
		},
	}
}

func fzf(sessions []models.Session, prompt string) *models.Session {
	index, err := fuzzy.FindMulti(
		sessions,
		func(i int) string {
			return sessions[i].Alias

		},
		fuzzy.WithHeader("Sessions - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			urlsStr := strings.Join(sessions[i].URLs, "\n\t")
			return fmt.Sprintf("Alias: %s \nURLs:\n%s \nCreated At: %s",
				sessions[i].Alias,
				urlsStr,
				sessions[i].CreatedAt,
			)
		}))

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &sessions[index[0]]
	}

	return nil
}
