package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Pgdb() (*sql.DB, error) {
	// Loading .env is a development convenience. Production configuration is
	// expected to come from the environment, so a missing file is not fatal.
	_ = godotenv.Load(".env")

	connectionStr := strings.TrimSpace(os.Getenv("DB_CONNECTION_STRING"))
	if connectionStr == "" {
		return nil, fmt.Errorf("DB_CONNECTION_STRING is not configured")
	}

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	fmt.Println("Database connected...")
	return conn, nil
}

func ExecuteQuery(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}
	return db.Exec(query, args...)
}
