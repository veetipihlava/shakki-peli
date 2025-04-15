package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_Rook_Valid_AllStraightDirections(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "R", Rank: 4, File: 4}, // Rd4
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"North - d4 to d8", "Rd4d8"},
		{"South - d4 to d1", "Rd4d1"},
		{"East - d4 to h4", "Rd4h4"},
		{"West - d4 to a4", "Rd4a4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid rook move %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Rook_InvalidDiagonalMove(t *testing.T) {
	pieces := []models.Piece{
		{ID: 2, GameID: 1, Color: true, Name: "R", Rank: 4, File: 4}, // Rd4
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Rd4f6", true, moves) // Diagonal
	if res.IsValidMove {
		t.Errorf("Expected invalid move: rook can't move diagonally, got: %+v", res)
	}
}

func Test_Rook_CaptureEnemyStraight(t *testing.T) {
	pieces := []models.Piece{
		{ID: 3, GameID: 1, Color: true, Name: "R", Rank: 4, File: 4},  // Rd4
		{ID: 4, GameID: 1, Color: false, Name: "P", Rank: 7, File: 4}, // enemy on d7
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Rd4d7", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected rook to capture enemy on d7, got: %+v", res)
	}
}

func Test_Rook_AttemptCaptureFriendly_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 5, GameID: 1, Color: true, Name: "R", Rank: 4, File: 4}, // Rd4
		{ID: 6, GameID: 1, Color: true, Name: "N", Rank: 6, File: 4}, // friendly on d6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Rd4d6", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: rook capturing friendly, got: %+v", res)
	}
}

func Test_Rook_CannotJumpOverPieces(t *testing.T) {
	pieces := []models.Piece{
		{ID: 7, GameID: 1, Color: true, Name: "R", Rank: 1, File: 1}, // Ra1
		{ID: 8, GameID: 1, Color: true, Name: "P", Rank: 2, File: 1}, // block a2
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Ra1a3", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: rook can't jump over pieces, got: %+v", res)
	}
}

func Test_Rook_CornerMove_Valid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 9, GameID: 1, Color: true, Name: "R", Rank: 1, File: 8}, // Rh1
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Rh1h5", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected valid rook move from corner, got: %+v", res)
	}
}

func Test_Rook_ZeroDistanceMove_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 10, GameID: 1, Color: true, Name: "R", Rank: 4, File: 4}, // Rd4
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Rd4d4", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: rook must move distance > 0, got: %+v", res)
	}
}
