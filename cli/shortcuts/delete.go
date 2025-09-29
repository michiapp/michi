package shortcuts

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(ctx context.Context, service *service.ShortcutService) *v2.Command {
	return &v2.Command{
		Name:    "delete",
		Usage:   "delete a shortcut",
		Aliases: []string{"d"},
		Action: func(c *v2.Context) error {

			shortcuts, err := service.GetAll(ctx)

			if err != nil {
				return err
			}

			selected := fzf(shortcuts, "Delete")

			if selected == nil {
				fmt.Println("No shortcut selected")
				return nil
			}

			err = service.DeleteFromAlias(ctx, selected.Alias)

			if err != nil {
				return err
			}

			fmt.Printf("Shortcut `%s` deleted.\n", selected.Alias)
			return nil
		},
	}
}
