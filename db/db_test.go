package db

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestPgdbRequiresConnectionString(t *testing.T) {
	t.Setenv("DB_CONNECTION_STRING", "")

	if _, err := Pgdb(); err == nil {
		t.Fatal("Pgdb() error = nil, want missing configuration error")
	}
}

func TestExecuteQueryRejectsNilDatabase(t *testing.T) {
	if _, err := ExecuteQuery(nil, "SELECT 1"); err == nil {
		t.Fatal("ExecuteQuery() error = nil, want nil database error")
	}
}

func TestPgdbIntegration(t *testing.T) {
	connectionString := os.Getenv("TEST_DB_CONNECTION_STRING")
	if connectionString == "" {
		t.Skip("set TEST_DB_CONNECTION_STRING to run the PostgreSQL integration test")
	}
	t.Setenv("DB_CONNECTION_STRING", connectionString)

	database, err := Pgdb()
	if err != nil {
		t.Fatalf("Pgdb() error = %v", err)
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		t.Fatalf("Ping() error = %v", err)
	}
}
