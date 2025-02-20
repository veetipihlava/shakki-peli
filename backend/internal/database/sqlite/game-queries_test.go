package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)
	require.NotZero(t, gameID)
}

func TestReadGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)
	require.NotZero(t, gameID)

	game, err := db.ReadGame(gameID)
	require.NoError(t, err)
	require.NotNil(t, game)
	require.Equal(t, gameID, game.ID)
	require.False(t, game.IsOver)
	require.NotZero(t, game.CreatedAt)
}
