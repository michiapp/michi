package history

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(service service.HistoryServiceIface) *v2.Command {
	lastFlag := &v2.IntFlag{
		Name:  "last",
		Usage: "purge the last (n) entries",
	}

	return &v2.Command{
		Name:  "delete",
		Usage: "delete an entry",
		Aliases: []string{"d"},
		Flags: []v2.Flag{
			lastFlag,
		},
		Action: func(ctx *v2.Context) error {
			last := ctx.Int("last")

			if last > 0 {
				history, err := service.GetRecentHistory(last)

				if err != nil {
					return err
				}

				for _, entry := range history {
					service.DeleteEntry(entry.ID)
				}

				fmt.Printf("Last (%d) have been successfully purged.\n", last)
				return nil
			}

			history, err := service.GetAllHistory()

			if err != nil {
				return err
			}

			selected := fzf(history, "Delete")

			if selected == nil {
				fmt.Println("No entry selected.")
				return nil
			}

			err = service.DeleteEntry(selected.ID)

			if err != nil {
				return err
			}

			fmt.Println("Deleted successfully.")
			return nil
		},
	}
}
