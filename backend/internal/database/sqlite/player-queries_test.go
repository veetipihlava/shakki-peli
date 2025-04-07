package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)

	user, err := db.CreateUser("TestPlayer")
	require.NoError(t, err)

	player, err := db.CreatePlayer(user.ID, game.ID)
	require.NoError(t, err)
	require.NotNil(t, player)
	require.Equal(t, user.ID, player.UserID)
	require.Equal(t, game.ID, player.GameID)
}

func TestReadPlayer(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)

	user, err := db.CreateUser("TestPlayer")
	require.NoError(t, err)

	player, err := db.CreatePlayer(user.ID, game.ID)
	require.NoError(t, err)

	readPlayer, err := db.ReadPlayer(user.ID, game.ID)
	require.NoError(t, err)
	require.NotNil(t, readPlayer)
	require.Equal(t, player.UserID, readPlayer.UserID)
	require.Equal(t, player.GameID, readPlayer.GameID)
	require.Equal(t, player.Color, readPlayer.Color)
}
