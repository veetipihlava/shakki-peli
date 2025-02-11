package chess

import (
	"github.com/veetipihlava/shakki-peli/internal/models"
)

// TODO tarviiko muuta tietoa
type Game struct {
	WhitePlayer models.Player
	BlackPlayer models.Player
	Pieces      []models.Piece
	History     []models.Move
}

func ValidateMove(playerID int64, move string, state Game) bool {
	return false
}
