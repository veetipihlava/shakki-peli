package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// The database implementation sqlite queries.
type Database struct {
	Connection *sql.DB
}

// Creates the database connection with libSQL.
func ConnectDatabase(tursoDatabaseURL string, tursoAuthToken string) (*Database, error) {
	url := fmt.Sprintf("%s?authToken=%s", tursoDatabaseURL, tursoAuthToken)
	connection, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	database := &Database{
		Connection: connection,
	}
	return database, nil
}
