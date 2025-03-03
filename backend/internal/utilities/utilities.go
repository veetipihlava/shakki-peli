package utilities

import (
	"errors"

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
func readChessGame(db *database.DatabaseService, gameID int64) (*chess.Game, error) {
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

// Processes the chess move and updates the database. Returns if the move is valid.
func ProcessChessMove(db *database.DatabaseService, playerID int64, gameID int64, move string) (chess.ValidationResult, error) {
	game, err := readChessGame(db, gameID)
	if err != nil {
		return chess.ValidationResult{}, err
	}
	if game == nil {
		return chess.ValidationResult{}, errors.New("game is null")
	}

	player, err := db.ReadPlayer(playerID, gameID)
	if err != nil {
		return chess.ValidationResult{}, err
	}
	if player == nil {
		return chess.ValidationResult{}, errors.New("player is null")
	}

	validationResult, piecesToUpdate := chess.ValidateMove(*game, move, player.Color)
	for _, pieceUpdate := range piecesToUpdate {
		if pieceUpdate.DeletePiece {
			err = db.DeletePiece(pieceUpdate.Piece.ID)
			if err != nil {
				return validationResult, err
			}
		} else {
			err = db.UpdatePiece(pieceUpdate.Piece)
			if err != nil {
				return validationResult, err
			}
		}
	}

	return validationResult, nil
}
