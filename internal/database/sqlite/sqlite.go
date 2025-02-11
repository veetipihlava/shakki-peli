package sqlite

import (
	"database/sql"
	"os"
	"path"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

// The database implementation for libsql.
type Database struct {
	Connection *sql.DB
}

// Initializes the test database. This is intended for testing purposes.
func InitializeTestDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, thisFile, _, _ := runtime.Caller(0)
	thisDirectory := path.Dir(thisFile)
	schemaPath := path.Join(thisDirectory, "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	database := &Database{
		Connection: db,
	}

	return database, nil
}
