package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMove(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)

	move, err := db.CreateMove(game.ID, "e4")
	require.NoError(t, err)
	require.NotNil(t, move)
	require.Equal(t, "e4", move.Notation)
	require.Equal(t, game.ID, move.GameID)
	require.NotZero(t, move.ID)
	require.NotZero(t, move.CreatedAt)
}

func TestReadMoves(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	game, err := db.CreateGame()
	require.NoError(t, err)

	_, err = db.CreateMove(game.ID, "e4")
	require.NoError(t, err)
	_, err = db.CreateMove(game.ID, "e5")
	require.NoError(t, err)

	moves, err := db.ReadMoves(game.ID)
	require.NoError(t, err)
	require.Len(t, moves, 2)
	require.Equal(t, "e4", moves[0].Notation)
}
