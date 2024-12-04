package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.io/ckshitij/go-service-template/config"
)

// NewPostgresDB initializes a connection to a PostgreSQL database
func NewPostgresDB(conf *config.Config) (*sql.DB, error) {
	// Build the PostgreSQL connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Databases["users"].Host,
		conf.Databases["users"].Port,
		conf.Databases["users"].User,
		conf.Databases["users"].Password,
		conf.Databases["users"].Database,
	)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
