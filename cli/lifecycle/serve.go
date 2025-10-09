package lifecycle

import (
	"github.com/OrbitalJin/michi/internal/server/manager"
	"github.com/urfave/cli/v2"
)

func Serve(sm *manager.ServerManager) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "detach",
				Usage: "run the server in background",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("detach") {
				err := sm.Daemonize()
				if err != nil {
					Doctor(sm, true).Action(ctx)
				}
				return sm.Daemonize()
			}

			err := sm.RunForeground()
			if err != nil {
				Doctor(sm, true).Action(ctx)
			}
			return sm.RunForeground()
		},
	}
}
