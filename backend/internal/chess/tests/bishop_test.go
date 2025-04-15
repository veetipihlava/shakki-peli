package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_Bishop_Valid_AllDiagonalDirections(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "B", Rank: 4, File: 4}, // Bd4
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"NE - d4 to f6", "Bd4f6"},
		{"NW - d4 to b6", "Bd4b6"},
		{"SE - d4 to f2", "Bd4f2"},
		{"SW - d4 to b2", "Bd4b2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid bishop move %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Bishop_InvalidStraightMove(t *testing.T) {
	pieces := []models.Piece{
		{ID: 2, GameID: 1, Color: true, Name: "B", Rank: 4, File: 4}, // Bd4
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Bd4d6", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: bishop can't move straight, got: %+v", res)
	}
}

func Test_Bishop_CannotJumpOverPieces(t *testing.T) {
	pieces := []models.Piece{
		{ID: 3, GameID: 1, Color: true, Name: "B", Rank: 1, File: 3}, // Bc1
		{ID: 4, GameID: 1, Color: true, Name: "P", Rank: 2, File: 4}, // Pd2 blocks path to e3
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Bc1e3", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: bishop can't jump over piece, got: %+v", res)
	}
}

func Test_Bishop_CaptureEnemyDiagonal(t *testing.T) {
	pieces := []models.Piece{
		{ID: 5, GameID: 1, Color: true, Name: "B", Rank: 4, File: 4},  // Bd4
		{ID: 6, GameID: 1, Color: false, Name: "P", Rank: 6, File: 6}, // enemy on f6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Bd4f6", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected valid bishop capture Bd4f6, got: %+v", res)
	}
}

func Test_Bishop_AttemptCaptureFriendly_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 7, GameID: 1, Color: true, Name: "B", Rank: 4, File: 4}, // Bd4
		{ID: 8, GameID: 1, Color: true, Name: "P", Rank: 6, File: 6}, // friendly on f6
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Bd4f6", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: bishop capturing friendly, got: %+v", res)
	}
}
