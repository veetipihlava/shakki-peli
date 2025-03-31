package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)
	require.NotNil(t, game)
	require.NotZero(t, game.ID)
	require.False(t, game.IsOver)
	require.NotZero(t, game.CreatedAt)
}

func TestReadGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)
	require.NotNil(t, game)

	readGame, err := db.ReadGame(game.ID)
	require.NoError(t, err)
	require.NotNil(t, readGame)
	require.Equal(t, game.ID, readGame.ID)
	require.False(t, readGame.IsOver)
	require.NotZero(t, readGame.CreatedAt)
}
