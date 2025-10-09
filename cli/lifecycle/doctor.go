package lifecycle

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/server/manager"
	"github.com/urfave/cli/v2"
)

func Doctor(sm *manager.ServerManager, bypass bool) *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "check michi status",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "fix",
				Usage: "remove stale PID file",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			fix := ctx.Bool("fix")

			ok, pid := sm.ValiatePID()

			if !ok {
				fmt.Printf("%s●%s Not running\n", internal.ColorRed, internal.ColorReset)
				return nil
			}

			isRunning := sm.ProcessExists(pid)

			if isRunning {
				fmt.Printf("%s●%s Running (PID: %d)\n",
					internal.ColorGreen, internal.ColorReset, pid)
				return nil
			}

			fmt.Printf("%s●%s Stale PID file found (PID: %d not running)\n",
				internal.ColorYellow, internal.ColorReset, pid)

			if fix || bypass {
				fmt.Printf("%s●%s Removing stale PID file\n",
					internal.ColorRed, internal.ColorReset)
				if err := sm.RemovePIDFile(); err != nil {
					return err
				}
				fmt.Printf("%s●%s Michi should be ready to run\n",
					internal.ColorGreen, internal.ColorReset)
			}

			return nil
		},
	}
}
