package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Pgdb() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", goDotEnvVariable("USER"), goDotEnvVariable("PASSWORD"), goDotEnvVariable("DBNAME"), goDotEnvVariable("SSL_MODE"))
	//fmt.Println(connectionStr)

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	rows, err := conn.Query("SELECT version();")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	rows.Close()

	conn.Close()
}
