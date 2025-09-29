package bangs

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(ctx context.Context, service *service.SPService) *v2.Command {
	return &v2.Command{
		Name:  "delete",
		Usage: "delete a bang",
		Action: func(_ *v2.Context) error {
			bangs, err := service.GetAll(ctx)

			if err != nil {
				return err
			}

			selected := fzf(bangs, "Delete")

			if selected == nil {
				fmt.Println("no bang selected")
				return nil
			}

			err = service.Delete(ctx, selected.ID)

			if err != nil {
				return err
			}

			fmt.Printf("Bang `%s` deleted.\n", selected.Tag)
			return nil

		},
	}
}
