package server

import (
	"database/sql"
	"fmt"

	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/router"
	"github.com/OrbitalJin/michi/internal/router/handler"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/sqlc"
	_ "modernc.org/sqlite"
)

type Server struct {
	queryParser parser.QueryParserIface
	router      router.RouterIface
	handler     *handler.Handler
	services    *service.Services
	queries     *sqlc.Queries
	config      *internal.Config
	conn        *sql.DB
}

func New(
	conn *sql.DB,
	config *internal.Config,
	serviceCgf *service.Config,
	bangParserCfg,
	shortcutParserCfg,
	seshParserCfg *parser.Config,
) (*Server, error) {
	q := sqlc.New(conn)

	qp, err := parser.NewQueryParser(
		bangParserCfg,
		shortcutParserCfg,
		seshParserCfg,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	psvc := service.NewSearchProviderService(qp.BangParser(), q, config.Service.DefaultProvider)

	hsvc := service.NewHistoryService(q)
	scsvc := service.NewShortcutService(q)
	seshsvc := service.NewSessionService(q)
	services := service.NewServices(psvc, hsvc, seshsvc, scsvc)

	handler := handler.NewHandler(config, qp, services, "q")

	router, err := router.NewRouter(handler)

	if err != nil {
		return nil, fmt.Errorf("failed to create router: %w", err)
	}

	router.Route()

	return &Server{
		queryParser: qp,
		queries:     q,
		services:    services,
		router:      router,
		handler:     handler,
		config:      config,
	}, nil
}

func (server *Server) GetServices() *service.Services {
	return server.services
}

func (server *Server) Serve() error {
	return server.router.Serve(server.config.Server.Port)
}

func (server *Server) GetConfig() *internal.Config {
	return server.config
}
