package shortcuts

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(service service.ShortcutServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "delete",
		Usage: "delete a shortcut",
		Aliases: []string{"d"},
		Action: func(c *v2.Context) error {

			shortcuts, err := service.GetAll()

			if err != nil {
				return err
			}

			selected := fzf(shortcuts, "Delete")

			if selected == nil {
				fmt.Println("No shortcut selected")
				return nil
			}

			err = service.Delete(selected.ID)

			if err != nil {
				return err
			}

			fmt.Printf("Shortcut `%s` deleted.\n", selected.Alias)
			return nil
		},
	}
}
