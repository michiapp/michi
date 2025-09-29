package shortcuts

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(ctx context.Context, service *service.ShortcutService) *v2.Command {
	return &v2.Command{
		Name:    "list",
		Usage:   "to list shortcuts",
		Aliases: []string{"l"},
		Action: func(c *v2.Context) error {
			shortcuts, err := service.GetAll(ctx)

			if err != nil {
				return err
			}

			selected := fzf(shortcuts, "Copy")

			if selected == nil {
				fmt.Println("No shortcut selected")
				return nil
			}

			clipboard.WriteAll(selected.Url)

			return nil
		},
	}
}
