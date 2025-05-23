package database

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/database/sqlite"
)

// Creates the database connection with libSQL in memory.
func CreateTestLibSQLConnection() (*DatabaseService, *sql.DB, error) {
	database, err := sqlite.InitializeTestDatabase()
	if err != nil {
		return nil, nil, err
	}

	db := &DatabaseService{
		Database: database,
	}

	return db, database.Connection, nil
}
