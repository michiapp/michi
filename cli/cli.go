package cli

import (
	"context"

	"github.com/OrbitalJin/michi/cli/bangs"
	"github.com/OrbitalJin/michi/cli/history"
	// "github.com/OrbitalJin/michi/cli/sessions"
	"github.com/OrbitalJin/michi/cli/lifecycle"
	"github.com/OrbitalJin/michi/cli/shortcuts"
	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/OrbitalJin/michi/internal/server/manager"
	v2 "github.com/urfave/cli/v2"
)

type Cli struct {
	server *server.Server
	cli    *v2.App
}

func New(server *server.Server) *v2.App {
	serverManager := manager.NewServerManager(server)
	ctx := context.Background()
	return &v2.App{
		Name:                 "michi",
		Usage:                "A super-charged search engine multiplexer 🚀",
		Version:              internal.Version,
		EnableBashCompletion: true,
		Commands: []*v2.Command{
			lifecycle.Serve(serverManager),
			lifecycle.Stop(serverManager),
			lifecycle.Doctor(serverManager),
			shortcuts.Root(ctx, server.GetServices().GetShortcutService()),
			history.Root(ctx, server.GetServices().GetHistoryService()),
			bangs.Root(ctx, server.GetServices().GetProvidersService()),
			// sessions.Root(server.GetServices().GetSessionService()),
		},
	}
}
