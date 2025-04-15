package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_Queen_Valid_AllDirections(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "Q", Rank: 4, File: 4}, // Qd4
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		// Diagonals
		{"NE - d4 to f6", "Qd4f6"},
		{"NW - d4 to b6", "Qd4b6"},
		{"SE - d4 to f2", "Qd4f2"},
		{"SW - d4 to b2", "Qd4b2"},
		// Straight
		{"North - d4 to d8", "Qd4d8"},
		{"South - d4 to d1", "Qd4d1"},
		{"East - d4 to h4", "Qd4h4"},
		{"West - d4 to a4", "Qd4a4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid queen move %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Queen_Invalid_NonLinearMove(t *testing.T) {
	pieces := []models.Piece{
		{ID: 2, GameID: 1, Color: true, Name: "Q", Rank: 4, File: 4}, // Qd4
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Qd4c6", true, moves) // Knight-like move
	if res.IsValidMove {
		t.Errorf("Expected invalid move: queen can't move like knight, got: %+v", res)
	}
}

func Test_Queen_CaptureEnemy(t *testing.T) {
	pieces := []models.Piece{
		{ID: 3, GameID: 1, Color: true, Name: "Q", Rank: 4, File: 4},
		{ID: 4, GameID: 1, Color: false, Name: "P", Rank: 6, File: 6}, // enemy f6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Qd4f6", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected valid queen capture on f6, got: %+v", res)
	}
}

func Test_Queen_AttemptCaptureFriendly_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 5, GameID: 1, Color: true, Name: "Q", Rank: 4, File: 4},
		{ID: 6, GameID: 1, Color: true, Name: "P", Rank: 6, File: 6}, // friendly on f6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Qd4f6", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: queen capturing friendly on f6, got: %+v", res)
	}
}

func Test_Queen_CannotJumpOverPieces(t *testing.T) {
	pieces := []models.Piece{
		{ID: 7, GameID: 1, Color: true, Name: "Q", Rank: 1, File: 1}, // Qa1
		{ID: 8, GameID: 1, Color: true, Name: "P", Rank: 2, File: 2}, // block b2
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Qa1c3", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: queen can't jump over b2, got: %+v", res)
	}
}

func Test_Queen_ZeroDistanceMove_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 9, GameID: 1, Color: true, Name: "Q", Rank: 4, File: 4},
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Qd4d4", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: queen must move distance > 0, got: %+v", res)
	}
}
