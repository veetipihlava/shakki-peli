package utilities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

func TestGetChessGameState(t *testing.T) {
	db, connection, err := database.CreateTestLibSQLConnection()
	require.NoError(t, err)
	defer connection.Close()

	// Create test players
	whitePlayerID, err := db.CreatePlayer("Alice", true)
	require.NoError(t, err)
	blackPlayerID, err := db.CreatePlayer("Bob", false)
	require.NoError(t, err)

	// Create a test game
	err = db.CreateGame(whitePlayerID, blackPlayerID)
	require.NoError(t, err)

	// Add some pieces to the game
	err = db.CreatePiece(1, true, "Knight", 1, 2)
	require.NoError(t, err)
	err = db.CreatePiece(1, false, "Bishop", 3, 4)
	require.NoError(t, err)

	// Add some moves to the game
	err = db.CreateMove(1, "e4")
	require.NoError(t, err)
	err = db.CreateMove(1, "e5")
	require.NoError(t, err)

	// Get the game state
	gameState, err := GetChessGameState(db, whitePlayerID, blackPlayerID)
	require.NoError(t, err)

	// Validate the game state
	require.Equal(t, whitePlayerID, gameState.WhitePlayer.ID)
	require.Equal(t, blackPlayerID, gameState.BlackPlayer.ID)
	require.Len(t, gameState.Pieces, 2)
	require.Len(t, gameState.History, 2)
	require.Equal(t, "Knight", gameState.Pieces[0].Name)
	require.Equal(t, "Bishop", gameState.Pieces[1].Name)
	require.Equal(t, "e4", gameState.History[0].Notation)
	require.Equal(t, "e5", gameState.History[1].Notation)
}
