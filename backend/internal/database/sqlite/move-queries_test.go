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

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	testMove := getTestMove(gameID, "e4")
	err = db.CreateMove(testMove.GameID, testMove.Notation)
	require.NoError(t, err)
}

func TestReadMoves(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	testMove := getTestMove(gameID, "e4")
	err = db.CreateMove(testMove.GameID, testMove.Notation)
	require.NoError(t, err)
	err = db.CreateMove(testMove.GameID, testMove.Notation)
	require.NoError(t, err)

	moves, err := db.ReadMoves(gameID)
	require.NoError(t, err)
	require.NotEmpty(t, moves)
	require.Len(t, moves, 2)
	require.Equal(t, testMove.Notation, moves[0].Notation)
}
