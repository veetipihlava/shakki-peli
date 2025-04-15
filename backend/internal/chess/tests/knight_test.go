package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_Knight_Valid_LMoves(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "N", Rank: 1, File: 2}, // Nb1
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"b1 to c3", "Nb1c3"},
		{"b1 to a3", "Nb1a3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid knight move %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Knight_Capture_EnemyPiece(t *testing.T) {
	pieces := []models.Piece{
		{ID: 2, GameID: 1, Color: true, Name: "N", Rank: 4, File: 4},  // Nd4
		{ID: 3, GameID: 1, Color: false, Name: "P", Rank: 6, File: 5}, // enemy on e6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Nd4e6", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected knight to capture enemy on e6, got: %+v", res)
	}
}

func Test_Knight_AttemptCapture_FriendlyPiece_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 4, GameID: 1, Color: true, Name: "N", Rank: 4, File: 4},
		{ID: 5, GameID: 1, Color: true, Name: "R", Rank: 6, File: 5}, // friendly on e6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Nd4e6", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: knight tries to capture friendly piece, got: %+v", res)
	}
}

func Test_Knight_InvalidMovePattern(t *testing.T) {
	pieces := []models.Piece{
		{ID: 6, GameID: 1, Color: true, Name: "N", Rank: 4, File: 4},
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Nd4d6", true, moves) // not L-shaped
	if res.IsValidMove {
		t.Errorf("Expected invalid move pattern for knight, got: %+v", res)
	}
}
