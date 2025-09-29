package history

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(service service.HistoryServiceIface) *v2.Command {
	allFlag := &v2.BoolFlag{
		Name:  "all",
		Usage: "list all history",
	}

	limitFlag := &v2.IntFlag{
		Name:  "limit",
		Usage: "limit of history",
	}

	return &v2.Command{
		Name:    "list",
		Usage:   "list history",
		Aliases: []string{"l"},
		Flags: []v2.Flag{
			allFlag,
			limitFlag,
		},
		Action: func(ctx *v2.Context) error {
			var history []models.SearchHistoryEvent
			var err error = nil

			all := ctx.Bool("all")
			limit := ctx.Int("limit")

			if all || limit < 1 {
				history, err = service.GetAllHistory()
			} else {
				history, err = service.GetRecentHistory(limit)
			}

			if err != nil {
				return err
			}

			selected := fzf(history, "Copy")

			if selected == nil {
				fmt.Println("No entry selected.")
				return nil
			}

			err = clipboard.WriteAll(selected.Query)

			if err != nil {
				return err
			}

			fmt.Printf("Selection copied to clipboard: %s\n", selected.Query)
			return nil
		},
	}
}
