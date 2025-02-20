package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testName string = "test-name"

func TestCreateUser(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	userID, err := db.CreateUser(testName)
	require.NoError(t, err)
	assert.NotZero(t, userID)
}

func TestReadUser(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	userID, err := db.CreateUser(testName)
	require.NoError(t, err)

	user, err := db.ReadUser(userID)
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, testName, user.Name)
}
