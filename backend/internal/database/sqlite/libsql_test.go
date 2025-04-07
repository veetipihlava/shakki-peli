package sqlite

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLibSQLConnection(t *testing.T) {
	err := godotenv.Load("../../../.env")
	require.NoError(t, err)

	tursoDatabaseURL := os.Getenv("TURSO_DATABASE_URL")
	tursoAuthToken := os.Getenv("TURSO_AUTH_TOKEN")

	db, err := ConnectDatabase(tursoDatabaseURL, tursoAuthToken)
	require.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Connection.Ping()
	assert.NoError(t, err)

	err = db.Connection.Close()
	assert.NoError(t, err)
}
