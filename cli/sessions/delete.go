package sessions

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(ctx context.Context, service *service.SessionService) *v2.Command {
	return &v2.Command{
		Name:  "delete",
		Usage: "delete a session",
		Aliases: []string{"d"},
		Action: func(c *v2.Context) error {
			sessions, err := service.GetSessionsWithUrls(ctx)

			if err != nil {
				return err
			}

			selected := fzf(sessions, "Delete")

			if selected == nil {
				fmt.Println("No session selected")
				return nil
			}

			err = service.Delete(ctx, selected.ID)

			if err != nil {
				return err
			}

			fmt.Printf("Session `%s` deleted.\n", selected.Alias)
			return nil
		},
	}
}
