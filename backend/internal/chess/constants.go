package chess

import "github.com/veetipihlava/shakki-peli/internal/models"

// Constant for readable code. The boolean for white player.
const White bool = true

// Constant for readable code. The boolean for black player.
const Black bool = false

// Get the initial chess board configuration.
func GetInitialChessGamePieces(gameID int64) []models.Piece {
	pieces := []models.Piece{
		{GameID: gameID, Color: White, Name: "Rook", Rank: 1, File: 1},
		{GameID: gameID, Color: White, Name: "Knight", Rank: 1, File: 2},
		{GameID: gameID, Color: White, Name: "Bishop", Rank: 1, File: 3},
		{GameID: gameID, Color: White, Name: "Queen", Rank: 1, File: 4},
		{GameID: gameID, Color: White, Name: "King", Rank: 1, File: 5},
		{GameID: gameID, Color: White, Name: "Bishop", Rank: 1, File: 6},
		{GameID: gameID, Color: White, Name: "Knight", Rank: 1, File: 7},
		{GameID: gameID, Color: White, Name: "Rook", Rank: 1, File: 8},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 1},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 2},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 3},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 4},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 5},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 6},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 7},
		{GameID: gameID, Color: White, Name: "Pawn", Rank: 2, File: 8},

		{GameID: gameID, Color: Black, Name: "Rook", Rank: 8, File: 1},
		{GameID: gameID, Color: Black, Name: "Knight", Rank: 8, File: 2},
		{GameID: gameID, Color: Black, Name: "Bishop", Rank: 8, File: 3},
		{GameID: gameID, Color: Black, Name: "Queen", Rank: 8, File: 4},
		{GameID: gameID, Color: Black, Name: "King", Rank: 8, File: 5},
		{GameID: gameID, Color: Black, Name: "Bishop", Rank: 8, File: 6},
		{GameID: gameID, Color: Black, Name: "Knight", Rank: 8, File: 7},
		{GameID: gameID, Color: Black, Name: "Rook", Rank: 8, File: 8},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 1},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 2},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 3},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 4},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 5},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 6},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 7},
		{GameID: gameID, Color: Black, Name: "Pawn", Rank: 7, File: 8},
	}

	return pieces
}
