package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_Pawn_MoveForward_OneStep_Valid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "P", Rank: 2, File: 5}, // Pe2
		{ID: 2, GameID: 1, Color: true, Name: "P", Rank: 2, File: 1}, // Pa2
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"e2 to e3", "Pe2e3"},
		{"a2 to a3", "Pa2a3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid move for %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Pawn_MoveForward_TwoStepsFromStart_Valid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 3, GameID: 1, Color: true, Name: "P", Rank: 2, File: 4}, // Pd2
		{ID: 4, GameID: 1, Color: true, Name: "P", Rank: 2, File: 2}, // Pb2
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"d2 to d4", "Pd2d4"},
		{"b2 to b4", "Pb2b4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid move for %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Pawn_MoveForward_TwoStepsFromNonStart_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 15, GameID: 1, Color: true, Name: "P", Rank: 3, File: 4}, // Pd3
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Pd3d5", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: pawn moved two steps from non-starting rank, got: %+v", res)
	}
}

func Test_Pawn_Capture_DiagonalEnemy_Valid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 5, GameID: 1, Color: true, Name: "P", Rank: 4, File: 4},  // Pd4
		{ID: 6, GameID: 1, Color: false, Name: "N", Rank: 5, File: 5}, // Black knight on e5
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Pd4e5", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected valid diagonal capture Pd4e5, got: %+v", res)
	}
}

func Test_Pawn_Capture_DiagonalFriendly_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 7, GameID: 1, Color: true, Name: "P", Rank: 4, File: 4},
		{ID: 8, GameID: 1, Color: true, Name: "N", Rank: 5, File: 5}, // White knight on e5
		{ID: 9, GameID: 1, Color: true, Name: "B", Rank: 5, File: 3}, // White bishop on c5
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"Attempt capture right to e5", "Pd4e5"},
		{"Attempt capture left to c5", "Pd4c5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if res.IsValidMove {
				t.Errorf("Expected invalid move when capturing friendly piece: %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_Pawn_MoveForward_EnemyBlocking_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 10, GameID: 1, Color: true, Name: "P", Rank: 4, File: 4},
		{ID: 11, GameID: 1, Color: false, Name: "P", Rank: 5, File: 4}, // Black pawn blocking
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Pd4d5", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: pawn blocked by enemy, got: %+v", res)
	}
}

func Test_Pawn_MoveForward_FriendlyBlocking_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 12, GameID: 1, Color: true, Name: "P", Rank: 4, File: 4},
		{ID: 13, GameID: 1, Color: true, Name: "N", Rank: 5, File: 4}, // Friendly knight in front
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Pd4d5", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected invalid move: pawn blocked by friendly, got: %+v", res)
	}
}

func Test_Pawn_Promotion_ReachesLastRank_TransformsToQueen(t *testing.T) {
	pieces := []models.Piece{
		{ID: 14, GameID: 1, Color: true, Name: "P", Rank: 7, File: 1},
	}
	moves := []models.Move{}

	res, updates := chess.ValidateMove(pieces, "Pa7a8", true, moves)
	if !res.IsValidMove {
		t.Fatalf("Expected promotion move to be valid, got: %+v", res)
	}

	promoted := false
	for _, u := range updates {
		if u.TransformPiece == "Q" && u.Piece.Rank == 8 {
			promoted = true
			break
		}
	}

	if !promoted {
		t.Errorf("Expected pawn to promote to Queen, but it did not. Updates: %+v", updates)
	}
}
