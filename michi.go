package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/OrbitalJin/michi/cli"
	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	goose.SetLogger(log.New(io.Discard, "", 0))
	goose.SetDialect(string(goose.DialectSQLite3))

	return goose.Up(db, "migrations")
}

func configureEnv() *internal.Config {
	isDebug := os.Getenv("ENV") == "dev"

	if isDebug {
		gin.SetMode(gin.DebugMode)
		log.Println("Running in development mode.")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	configDir, err := internal.SetupConfigDir()
	if err != nil {
		log.Fatalf("Failed to prepare configuration directory: %v", err)
	}

	configPath := filepath.Join(configDir, "config.yaml")
	config, err := internal.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	if err = internal.SetupHydrationFile(); err != nil {
		log.Fatalf("Failed to hydrate database: %v", err)
	}

	return config
}

func main() {
	config := configureEnv()
	bangParserConfig := parser.NewConfig(config.Parser.BangPrefix)
	shortcutParserConfig := parser.NewConfig(config.Parser.ShortcutPrefix)
	sessionParserConfig := parser.NewConfig(config.Parser.SessionPrefix)
	serviceConfig := service.NewConfig(config.Service.KeepTrack, config.Service.DefaultProvider)

	conn, err := sql.Open("sqlite", config.DBPath)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer conn.Close()

	if err := migrate(conn); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	michi, err := server.New(
		conn,
		config,
		serviceConfig,
		bangParserConfig,
		shortcutParserConfig,
		sessionParserConfig,
	)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	app := cli.New(michi)
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%s●%s %v\n", internal.ColorRed, internal.ColorReset, err)
	}
}
