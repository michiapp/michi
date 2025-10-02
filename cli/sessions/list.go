package sessions

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(ctx context.Context, service *service.SessionService) *v2.Command {
	return &v2.Command{
		Name:  "list",
		Usage: "list sessions",
		Aliases: []string{"l"},
		Action: func(_ *v2.Context) error {
			sessions, err := service.GetSessionsWithUrls(ctx)
			if err != nil {
				return err
			}

			selected := fzf(sessions, "Copy")

			if selected == nil {
				fmt.Println("no session selected")
				return nil
			}

			clipboard.WriteAll(selected.Alias)

			return nil
		},
	}
}
