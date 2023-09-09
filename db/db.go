package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/halosatrio/bebop/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDB(db *sql.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS bebop_user (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
    );`

	_, err := db.Exec(schema)
	return err
}

func InitDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.GetEnv("DB_HOST", ""),
		config.GetEnv("DB_PORT", ""),
		config.GetEnv("DB_USER", ""),
		config.GetEnv("DB_NAME", ""),
		config.GetEnv("DB_PASSWORD", ""),
	)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
