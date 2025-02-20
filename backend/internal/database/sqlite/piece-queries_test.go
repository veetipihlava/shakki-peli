package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func TestCreatePieces(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	pieces := chess.GetInitialChessGamePieces(gameID)

	err = db.CreatePieces(pieces)
	require.NoError(t, err)
}

func TestReadPieces(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	initialPieces := chess.GetInitialChessGamePieces(gameID)

	err = db.CreatePieces(initialPieces)
	require.NoError(t, err)

	databasePieces, err := db.ReadPieces(gameID)
	require.NoError(t, err)
	require.NotEmpty(t, databasePieces)
	require.Equal(t, len(initialPieces), len(databasePieces))
}

func TestUpdatePiece(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	initialPieces := chess.GetInitialChessGamePieces(gameID)

	err = db.CreatePieces(initialPieces)
	require.NoError(t, err)

	pieces, err := db.ReadPieces(gameID)
	require.NoError(t, err)

	updatedPiece := models.Piece{
		ID:     pieces[0].ID,
		GameID: gameID,
		Color:  true,
		Name:   "horse",
		Rank:   3,
		File:   4,
	}
	err = db.UpdatePiece(updatedPiece)
	require.NoError(t, err)

	updatedPieces, err := db.ReadPieces(gameID)
	require.NoError(t, err)
	require.Equal(t, updatedPiece.ID, updatedPieces[0].ID)
	require.Equal(t, updatedPiece.GameID, updatedPieces[0].GameID)
	require.Equal(t, updatedPiece.Color, updatedPieces[0].Color)
	require.Equal(t, updatedPiece.Name, updatedPieces[0].Name)
	require.Equal(t, updatedPiece.Rank, updatedPieces[0].Rank)
	require.Equal(t, updatedPiece.File, updatedPieces[0].File)
}

func TestDeletePiece(t *testing.T) {
	db, err := InitializeTestDatabase()
	require.NoError(t, err)
	defer db.Connection.Close()

	gameID, err := db.CreateGame()
	require.NoError(t, err)

	initialPieces := chess.GetInitialChessGamePieces(gameID)

	err = db.CreatePieces(initialPieces)
	require.NoError(t, err)

	pieces, err := db.ReadPieces(gameID)
	require.NoError(t, err)

	err = db.DeletePiece(pieces[0].ID)
	require.NoError(t, err)

	updatedPieces, err := db.ReadPieces(gameID)
	require.NoError(t, err)
	require.Equal(t, len(pieces)-1, len(updatedPieces))
}
