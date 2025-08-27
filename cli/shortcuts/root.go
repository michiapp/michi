package shortcuts

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(service service.ShortcutServiceIface) *v2.Command {
	return &v2.Command{
		Name:    "shortcuts",
		Usage:   "to manage shortcuts",
		Aliases: []string{"sc"},
		Subcommands: []*v2.Command{
			list(service),
			add(service),
			delete(service),
		},
	}
}

func fzf(shortcuts []models.Shortcut, prompt string) *models.Shortcut {
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
				shortcuts[i].URL,
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
