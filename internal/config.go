package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/OrbitalJin/michi/public"
	"gopkg.in/yaml.v3"
)

var Version = "dev"

type UserConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Parser struct {
		BangPrefix     string `yaml:"bang_prefix"`
		ShortcutPrefix string `yaml:"shortcut_prefix"`
		SessionPrefix  string `yaml:"session_prefix"`
	} `yaml:"parser"`

	Service struct {
		KeepTrack       bool   `yaml:"keep_track"`
		DefaultProvider string `yaml:"default_provider"`
	} `yaml:"service"`
}

// Full runtime config (includes internal-only fields)
type Config struct {
	UserConfig `yaml:",inline"`

	PidFile string `yaml:"-"` // internal only
	LogFile string `yaml:"-"` // internal only
	DBPath  string `yaml:"-"` // internal only
}

func NewDefaultAppConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Warning: Could not get user home directory: %v. Using /tmp paths.", err)
		homeDir = "/tmp"
	}

	cfg := &Config{}

	cfg.Server.Port = ":5980"
	cfg.Parser.BangPrefix = "!"
	cfg.Parser.ShortcutPrefix = "@"
	cfg.Parser.SessionPrefix = "#"
	cfg.Service.KeepTrack = true
	cfg.Service.DefaultProvider = "g"

	cfg.PidFile = filepath.Join(homeDir, ".michi", "michi.proc.pid")
	cfg.LogFile = filepath.Join(homeDir, ".michi", "michi.log")
	cfg.DBPath = filepath.Join(homeDir, ".michi", "michi.db")

	return cfg
}

func LoadConfig(configFilePath string) (*Config, error) {
	cfg := NewDefaultAppConfig()

	if configFilePath == "" {
		log.Println("No config file path provided. Using defaults.")
		return cfg, nil
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Config not found at '%s'. Creating default config.", configFilePath)
			return createDefaultConfigFile(configFilePath, cfg)
		}
		return nil, fmt.Errorf("failed to read config file '%s': %w", configFilePath, err)
	}

	if err := yaml.Unmarshal(data, &cfg.UserConfig); err != nil {
		return nil, fmt.Errorf("failed to parse config file '%s': %w", configFilePath, err)
	}

	return cfg, nil
}

func SetupConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".michi")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return "", err
	}
	return configDir, nil
}

func SetupHydrationFile() error {
	configDir, err := SetupConfigDir()
	if err != nil {
		return err
	}

	dbPath := filepath.Join(configDir, "michi.db")

	if _, err := os.Stat(dbPath); err == nil {
		return nil

	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	data, err := public.Content.ReadFile("assets/hydrator.db")
	if err != nil {
		return err
	}

	return os.WriteFile(dbPath, data, 0o644)
}

func createDefaultConfigFile(path string, cfg *Config) (*Config, error) {
	out, err := yaml.Marshal(cfg.UserConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal default config: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(path, out, 0o644); err != nil {
		return nil, fmt.Errorf("failed to write default config file '%s': %w", path, err)
	}

	return cfg, nil
}
