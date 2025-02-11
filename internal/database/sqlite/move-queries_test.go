package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func getTestMove(gameID int64, notation string) models.Move {
	return models.Move{
		GameID:   gameID,
		Notation: notation,
	}
}

func TestCreateMove(t *testing.T) {
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

	testMove := getTestMove(game.ID, "e4")
	err = db.CreateMove(testMove.GameID, testMove.Notation)
	require.NoError(t, err)
}

func TestReadMoves(t *testing.T) {
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

	testMove := getTestMove(game.ID, "e4")
	err = db.CreateMove(testMove.GameID, testMove.Notation)
	require.NoError(t, err)

	moves, err := db.ReadMoves(game.ID)
	require.NoError(t, err)
	require.NotEmpty(t, moves)
	require.Equal(t, testMove.Notation, moves[0].Notation)
}
