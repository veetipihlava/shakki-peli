package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func getTestPiece(gameID int64, color bool, name string, rank, file int) models.Piece {
	return models.Piece{
		GameID: gameID,
		Color:  color,
		Name:   name,
		Rank:   rank,
		File:   file,
	}
}

func TestCreatePiece(t *testing.T) {
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

	testPiece := getTestPiece(game.ID, true, "Knight", 1, 2)
	err = db.CreatePiece(testPiece.GameID, testPiece.Color, testPiece.Name, testPiece.Rank, testPiece.File)
	require.NoError(t, err)
}

func TestReadPieces(t *testing.T) {
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

	testPiece := getTestPiece(game.ID, true, "Knight", 1, 2)
	err = db.CreatePiece(testPiece.GameID, testPiece.Color, testPiece.Name, testPiece.Rank, testPiece.File)
	require.NoError(t, err)

	pieces, err := db.ReadPieces(game.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pieces)
	require.Equal(t, testPiece.Name, pieces[0].Name)
}

func TestUpdatePiece(t *testing.T) {
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

	testPiece := getTestPiece(game.ID, true, "Knight", 1, 2)
	err = db.CreatePiece(testPiece.GameID, testPiece.Color, testPiece.Name, testPiece.Rank, testPiece.File)
	require.NoError(t, err)

	pieces, err := db.ReadPieces(game.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pieces)

	err = db.UpdatePiece(pieces[0].ID, 3, 4)
	require.NoError(t, err)

	updatedPieces, err := db.ReadPieces(game.ID)
	require.NoError(t, err)
	require.Equal(t, 3, updatedPieces[0].Rank)
	require.Equal(t, 4, updatedPieces[0].File)
}

func TestDeletePiece(t *testing.T) {
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

	testPiece := getTestPiece(game.ID, true, "Knight", 1, 2)
	err = db.CreatePiece(testPiece.GameID, testPiece.Color, testPiece.Name, testPiece.Rank, testPiece.File)
	require.NoError(t, err)

	pieces, err := db.ReadPieces(game.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pieces)

	err = db.DeletePiece(pieces[0].ID)
	require.NoError(t, err)

	deletedPieces, err := db.ReadPieces(game.ID)
	require.NoError(t, err)
	require.Empty(t, deletedPieces)
}
