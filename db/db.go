package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Pgdb() (*sql.DB, error) {
	err := godotenv.Load(".env") // "/etc/secrets/.env"
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connectionStr := goDotEnvVariable("DB_CONNECTION_STRING")
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
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
