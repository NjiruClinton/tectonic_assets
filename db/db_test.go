package db

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestPgdb(t *testing.T) {
	// Set up environment variable for testing
	os.Setenv("DB_CONNECTION_STRING", "your_test_connection_string")

	// Test database connection
	db, err := Pgdb()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Check if the connection is valid
	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}
}

func TestExecuteQuery(t *testing.T) {
	// Set up environment variable for testing
	os.Setenv("DB_CONNECTION_STRING", "your_test_connection_string")

	// Test database connection
	db, err := Pgdb()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test query execution
	query := "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name VARCHAR(50))"
	_, err = ExecuteQuery(db, query)
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}

	// Clean up
	_, err = ExecuteQuery(db, "DROP TABLE test_table")
	if err != nil {
		t.Fatalf("Failed to drop test table: %v", err)
	}
}
