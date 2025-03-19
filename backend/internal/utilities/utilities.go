package utilities

import (
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

// Creates chess game and returns the white player, black player and
func CreateNewChessGame(db *database.DatabaseService) (int64, error) {
	gameID, err := db.CreateGame()
	if err != nil {
		return 0, err
	}

	initialPieces := chess.GetInitialChessGamePieces(gameID)
	err = db.CreatePieces(initialPieces)
	if err != nil {
		return 0, err
	}

	return gameID, nil
}

// Reads chess game from database.
func ReadChessGame(db *database.DatabaseService, gameID int64) (*chess.Game, error) {
	pieces, err := db.ReadPieces(gameID)
	if err != nil {
		return nil, err
	}

	moves, err := db.ReadMoves(gameID)
	if err != nil {
		return nil, err
	}

	chessGame := &chess.Game{
		Pieces: pieces,
		Moves:  moves,
	}

	return chessGame, nil
}
