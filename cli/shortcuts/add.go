package shortcuts

import (
	"context"
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	v2 "github.com/urfave/cli/v2"
)

func add(ctx context.Context, service *service.ShortcutService) *v2.Command {
	aliasFlag := &v2.StringFlag{
		Name:     "alias",
		Usage:    "alias for the shortcut",
		Required: true,
	}

	urlFlag := &v2.StringFlag{
		Name:     "url",
		Usage:    "url for the shortcut that the alias points to",
		Required: true,
	}

	return &v2.Command{
		Name:    "add",
		Usage:   "add a shortcut",
		Aliases: []string{"a"},
		Flags: []v2.Flag{
			aliasFlag,
			urlFlag,
		},
		Action: func(c *v2.Context) error {
			alias := c.String("alias")
			url := c.String("url")

			err := service.Insert(ctx, sqlc.InsertShortcutParams{
				Alias: alias,
				Url:   url,
			})

			if err != nil {
				return err
			}

			fmt.Printf("Successfully added shortcut `%s` to `%s`\n", alias, url)
			return nil
		},
	}
}
