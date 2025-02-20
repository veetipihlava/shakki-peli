package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/chess"
)

func TestCreatePlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	userID, err := db.CreateUser(testName)
	require.NoError(t, err)

	err = db.CreatePlayer(gameID, userID, chess.White)
	require.NoError(t, err)
}

func TestReadPlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	userID, err := db.CreateUser(testName)
	require.NoError(t, err)

	err = db.CreatePlayer(gameID, userID, chess.White)
	require.NoError(t, err)

	player, err := db.ReadPlayer(gameID, userID)
	require.NoError(t, err)
	assert.NotNil(t, player)
	assert.Equal(t, gameID, player.GameID)
	assert.Equal(t, userID, player.UserID)
	assert.Equal(t, chess.White, player.Color)
}
