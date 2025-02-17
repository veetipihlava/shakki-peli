package utilities

import (
	"errors"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

func GetChessGameState(db *database.DatabaseService, firstPlayer int64, secondPlayer int64) (chess.Game, error) {
	errorGame := chess.Game{}

	game, err := db.ReadGame(firstPlayer, secondPlayer)
	if err != nil {
		return errorGame, err
	}
	if game == nil {
		return errorGame, errors.New("no game exists for these players")
	}

	whitePlayer, err := db.ReadPlayer(game.WhitePlayerID)
	if err != nil {
		return errorGame, err
	}
	if whitePlayer == nil {
		return errorGame, errors.New("no white player exists")
	}

	blackPlayer, err := db.ReadPlayer(game.BlackPlayerID)
	if err != nil {
		return errorGame, err
	}
	if blackPlayer == nil {
		return errorGame, errors.New("no black player exists")
	}

	pieces, err := db.ReadPieces(game.ID)
	if err != nil {
		return errorGame, err
	}

	moves, err := db.ReadMoves(game.ID)
	if err != nil {
		return errorGame, err
	}

	chessGame := chess.Game{
		WhitePlayer: *whitePlayer,
		BlackPlayer: *blackPlayer,
		Pieces:      pieces,
		History:     moves,
	}

	return chessGame, nil
}
