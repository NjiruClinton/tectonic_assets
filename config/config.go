package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const defaultProfileInterval = 5 * time.Second

// Config contains the small set of settings needed to run the development
// dashboard and, optionally, monitor a local process.
type Config struct {
	ServerAddress     string
	TargetName        string
	TargetProcessPath string
	ProfileInterval   time.Duration
}

// Load reads configuration from the environment. A .env file is convenient in
// development, but is deliberately optional so the server can run in other
// environments without one.
func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		ServerAddress:     valueOrDefault("SERVER_ADDRESS", ":8080"),
		TargetName:        strings.TrimSpace(os.Getenv("TARGET_NAME")),
		TargetProcessPath: strings.TrimSpace(os.Getenv("TARGET_PROCESS_PATH")),
		ProfileInterval:   defaultProfileInterval,
	}

	if interval := strings.TrimSpace(os.Getenv("PROFILE_INTERVAL")); interval != "" {
		parsed, err := time.ParseDuration(interval)
		if err != nil {
			return Config{}, fmt.Errorf("parse PROFILE_INTERVAL: %w", err)
		}
		if parsed <= 0 {
			return Config{}, fmt.Errorf("PROFILE_INTERVAL must be greater than zero")
		}
		cfg.ProfileInterval = parsed
	}

	if cfg.TargetProcessPath != "" {
		absolutePath, err := filepath.Abs(cfg.TargetProcessPath)
		if err != nil {
			return Config{}, fmt.Errorf("resolve TARGET_PROCESS_PATH: %w", err)
		}
		cfg.TargetProcessPath = filepath.Clean(absolutePath)
		if cfg.TargetName == "" {
			cfg.TargetName = filepath.Base(cfg.TargetProcessPath)
		}
	}

	return cfg, nil
}

func valueOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
