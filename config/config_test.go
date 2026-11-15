package config

import (
	"path/filepath"
	"testing"
	"time"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("SERVER_ADDRESS", "")
	t.Setenv("TARGET_NAME", "")
	t.Setenv("TARGET_PROCESS_PATH", "")
	t.Setenv("PROFILE_INTERVAL", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.ServerAddress != ":8080" {
		t.Fatalf("ServerAddress = %q, want :8080", cfg.ServerAddress)
	}
	if cfg.ProfileInterval != 5*time.Second {
		t.Fatalf("ProfileInterval = %v, want 5s", cfg.ProfileInterval)
	}
}

func TestLoadTargetAndInterval(t *testing.T) {
	t.Setenv("TARGET_NAME", "")
	t.Setenv("TARGET_PROCESS_PATH", "./tmp/example")
	t.Setenv("PROFILE_INTERVAL", "250ms")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if !filepath.IsAbs(cfg.TargetProcessPath) {
		t.Fatalf("TargetProcessPath = %q, want absolute path", cfg.TargetProcessPath)
	}
	if cfg.TargetName != "example" {
		t.Fatalf("TargetName = %q, want example", cfg.TargetName)
	}
	if cfg.ProfileInterval != 250*time.Millisecond {
		t.Fatalf("ProfileInterval = %v, want 250ms", cfg.ProfileInterval)
	}
}

func TestLoadRejectsInvalidInterval(t *testing.T) {
	t.Setenv("PROFILE_INTERVAL", "not-a-duration")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want invalid interval error")
	}
}
