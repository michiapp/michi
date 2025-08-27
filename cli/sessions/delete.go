package sessions

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(service service.SessionServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "delete",
		Usage: "delete a session",
		Aliases: []string{"d"},
		Action: func(c *v2.Context) error {
			sessions, err := service.GetAll()

			if err != nil {
				return err
			}

			selected := fzf(sessions, "Delete")

			if selected == nil {
				fmt.Println("No session selected")
				return nil
			}

			err = service.Delete(selected.ID)

			if err != nil {
				return err
			}

			fmt.Printf("Session `%s` deleted.\n", selected.Alias)
			return nil
		},
	}
}
