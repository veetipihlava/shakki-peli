package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTestLibSQLConnection(t *testing.T) {
	db, connection, err := CreateTestLibSQLConnection()
	require.NoError(t, err)
	assert.NotNil(t, db)
	assert.NotNil(t, connection)

	err = connection.Ping()
	assert.NoError(t, err)

	err = connection.Close()
	assert.NoError(t, err)
}
