package bangs

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(ctx context.Context, service *service.SPService) *v2.Command {
	return &v2.Command{
		Name:  "list",
		Usage: "list bangs",
		Action: func(_ *v2.Context) error {
			bangs, err := service.GetAll(ctx)

			if err != nil {
				return err
			}

			selected := fzf(bangs, "Copy")

			if selected == nil {
				fmt.Println("no bang selected")
				return nil
			}

			clipboard.WriteAll(selected.Tag)
			fmt.Printf("Bang `%s` copied.\n", selected.Tag)
			return nil
		},
	}
}
