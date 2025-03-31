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

	user, err := db.CreateUser(testName)
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testName, user.Name)
	assert.NotZero(t, user.ID)
}

func TestReadUser(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	user, err := db.CreateUser(testName)
	require.NoError(t, err)

	readUser, err := db.ReadUser(user.ID)
	require.NoError(t, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, user.ID, readUser.ID)
	assert.Equal(t, user.Name, readUser.Name)
}
