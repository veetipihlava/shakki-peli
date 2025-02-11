package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	whitePlayer := getTestPlayer("Alice", true)
	blackPlayer := getTestPlayer("Bob", false)
	whitePlayerID, err := db.CreatePlayer(whitePlayer.Name, whitePlayer.Color)
	require.NoError(t, err)
	blackPlayerID, err := db.CreatePlayer(blackPlayer.Name, blackPlayer.Color)
	require.NoError(t, err)

	err = db.CreateGame(whitePlayerID, blackPlayerID)
	require.NoError(t, err)
}

func TestReadGame(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	whitePlayer := getTestPlayer("Alice", true)
	blackPlayer := getTestPlayer("Bob", false)
	whitePlayerID, err := db.CreatePlayer(whitePlayer.Name, whitePlayer.Color)
	require.NoError(t, err)
	blackPlayerID, err := db.CreatePlayer(blackPlayer.Name, blackPlayer.Color)
	require.NoError(t, err)

	err = db.CreateGame(whitePlayerID, blackPlayerID)
	require.NoError(t, err)

	game, err := db.ReadGame(whitePlayerID, blackPlayerID)
	require.NoError(t, err)
	require.NotNil(t, game)
	require.Equal(t, whitePlayerID, game.WhitePlayerID)
	require.Equal(t, blackPlayerID, game.BlackPlayerID)
}
