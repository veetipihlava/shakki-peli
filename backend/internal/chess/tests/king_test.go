package chess_test

import (
	"testing"

	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func Test_King_BasicMovement_Valid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: true, Name: "K", Rank: 4, File: 4},
	}
	moves := []models.Move{}

	tests := []struct {
		name string
		move string
	}{
		{"Move up", "Kd4d5"},
		{"Move down", "Kd4d3"},
		{"Move left", "Kd4c4"},
		{"Move right", "Kd4e4"},
		{"Move NE", "Kd4e5"},
		{"Move NW", "Kd4c5"},
		{"Move SE", "Kd4e3"},
		{"Move SW", "Kd4c3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := chess.ValidateMove(pieces, tt.move, true, moves)
			if !res.IsValidMove {
				t.Errorf("Expected valid move %s, got: %+v", tt.move, res)
			}
		})
	}
}

func Test_King_CaptureEnemy(t *testing.T) {
	pieces := []models.Piece{
		{ID: 2, GameID: 1, Color: true, Name: "K", Rank: 4, File: 4},
		{ID: 3, GameID: 1, Color: false, Name: "P", Rank: 5, File: 5},
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Kd4e5", true, moves)
	if !res.IsValidMove {
		t.Errorf("Expected valid capture move, got: %+v", res)
	}
}

func Test_King_MoveIntoCheck_Invalid(t *testing.T) {
	pieces := []models.Piece{
		{ID: 6, GameID: 1, Color: true, Name: "K", Rank: 4, File: 4},
		{ID: 7, GameID: 1, Color: false, Name: "B", Rank: 6, File: 6},
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Kd4e5", true, moves)
	if res.IsValidMove {
		t.Errorf("Expected king not to be able to move into check. Got: %+v", res)
	}
}

func Test_Bishop_Move_Causes_KingInCheck(t *testing.T) {
	pieces := []models.Piece{
		{ID: 8, GameID: 1, Color: false, Name: "K", Rank: 1, File: 1}, // Ka1
		{ID: 9, GameID: 1, Color: true, Name: "B", Rank: 3, File: 3},  // Black bishop on c3
		{ID: 10, GameID: 1, Color: true, Name: "P", Rank: 2, File: 2}, // Blocking pawn b2
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Pb2b3", true, moves)
	if !res.KingInCheck {
		t.Errorf("Expected KingInCheck to be true after move clears diagonal. Got: %+v", res)
	}
}

func Test_Bishop_Capture_Causes_KingInCheck(t *testing.T) {
	pieces := []models.Piece{
		{ID: 11, GameID: 1, Color: false, Name: "K", Rank: 1, File: 1}, // Ka1
		{ID: 12, GameID: 1, Color: true, Name: "B", Rank: 3, File: 3},  // Bc3
		{ID: 13, GameID: 1, Color: false, Name: "P", Rank: 2, File: 2}, // Pb2
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Bc3b2", true, moves)
	if !res.KingInCheck {
		t.Errorf("Expected KingInCheck to be true after bishop captures and exposes king. Got: %+v", res)
	}
}

func Test_King_IsCheckmate_GameOverTriggered(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: false, Name: "K", Rank: 1, File: 1}, // Black king a1
		{ID: 2, GameID: 1, Color: true, Name: "R", Rank: 2, File: 2},  // White rook on b2
		{ID: 3, GameID: 1, Color: true, Name: "K", Rank: 3, File: 1},  // White king on a3
		{ID: 4, GameID: 1, Color: false, Name: "P", Rank: 2, File: 1}, // Black pawn blocks a2
		{ID: 5, GameID: 1, Color: false, Name: "P", Rank: 1, File: 2}, // Black pawn blocks b1
	}
	moves := []models.Move{}

	// Rook moves to a2, putting king in check — king cannot escape, it's checkmate
	res, _ := chess.ValidateMove(pieces, "Rb2a2", true, moves)

	if !res.IsValidMove {
		t.Errorf("Expected move to be valid. Got: %+v", res)
	}
	if !res.GameOver.Checkmate {
		t.Errorf("Expected Checkmate to be true. Got: %+v", res.GameOver)
	}
}

func Test_King_IsCheckmate_WithFriendlyBlockers(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: false, Name: "K", Rank: 1, File: 2}, // Black king on b1
		{ID: 2, GameID: 1, Color: false, Name: "P", Rank: 1, File: 1}, // Friendly pawn on a1
		{ID: 3, GameID: 1, Color: false, Name: "P", Rank: 1, File: 3}, // Friendly pawn on c1
		{ID: 4, GameID: 1, Color: true, Name: "R", Rank: 3, File: 2},  // White rook on b3 (about to move)
		{ID: 5, GameID: 1, Color: true, Name: "K", Rank: 3, File: 3},  // White king on c3 (protects rook)
	}
	moves := []models.Move{}

	// Move rook from b3 to b2, delivering checkmate
	res, _ := chess.ValidateMove(pieces, "Rb3b2", true, moves)

	if !res.IsValidMove {
		t.Errorf("Expected move to be valid. Got: %+v", res)
	}
	if !res.GameOver.Checkmate {
		t.Errorf("Expected Checkmate to be true. Got: %+v", res.GameOver)
	}
}

func Test_King_EscapesCheck_ByBlockingPiece(t *testing.T) {
	pieces := []models.Piece{
		{ID: 1, GameID: 1, Color: false, Name: "K", Rank: 1, File: 1}, // Black king a1
		{ID: 2, GameID: 1, Color: true, Name: "R", Rank: 1, File: 8},  // White rook h1 (checking along rank)
		{ID: 3, GameID: 1, Color: false, Name: "N", Rank: 2, File: 1}, // Black knight a2 (can block on c1)
	}
	moves := []models.Move{}

	res, _ := chess.ValidateMove(pieces, "Na2c1", false, moves)

	if !res.IsValidMove {
		t.Errorf("Expected move to be valid. Got: %+v", res)
	}
	if res.GameOver.Checkmate {
		t.Errorf("Expected Checkmate to be false. Got: %+v", res.GameOver)
	}
}
