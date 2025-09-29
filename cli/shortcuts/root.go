package shortcuts

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(ctx context.Context, service *service.ShortcutService) *v2.Command {
	return &v2.Command{
		Name:    "shortcuts",
		Usage:   "to manage shortcuts",
		Aliases: []string{"sc"},
		Subcommands: []*v2.Command{
			list(ctx, service),
			add(ctx, service),
			delete(ctx, service),
		},
	}
}

func fzf(shortcuts []sqlc.Shortcut, prompt string) *sqlc.Shortcut {
	index, err := fuzzy.FindMulti(
		shortcuts,
		func(i int) string {
			return shortcuts[i].Alias

		},
		fuzzy.WithHeader("Shortcuts - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Alias: %s \nURL: (%s) \nCreated At: %s",
				shortcuts[i].Alias,
				shortcuts[i].Url,
				shortcuts[i].CreatedAt,
			)
		}))

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &shortcuts[index[0]]
	}

	return nil
}
