package chess

import "github.com/veetipihlava/shakki-peli/internal/models"

// The boolean for white player.
const White bool = true

// Constant for readable code. The boolean for black player.
const Black bool = false

// Get the initial chess board configuration.
func GetInitialChessGamePieces(gameID int64) []models.Piece {
	pieces := []models.Piece{
		{GameID: gameID, Color: White, Name: "R", Rank: 1, File: 1},
		{GameID: gameID, Color: White, Name: "N", Rank: 1, File: 2},
		{GameID: gameID, Color: White, Name: "B", Rank: 1, File: 3},
		{GameID: gameID, Color: White, Name: "Q", Rank: 1, File: 4},
		{GameID: gameID, Color: White, Name: "K", Rank: 1, File: 5},
		{GameID: gameID, Color: White, Name: "B", Rank: 1, File: 6},
		{GameID: gameID, Color: White, Name: "N", Rank: 1, File: 7},
		{GameID: gameID, Color: White, Name: "R", Rank: 1, File: 8},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 1},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 2},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 3},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 4},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 5},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 6},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 7},
		{GameID: gameID, Color: White, Name: "P", Rank: 2, File: 8},

		{GameID: gameID, Color: Black, Name: "R", Rank: 8, File: 1},
		{GameID: gameID, Color: Black, Name: "K", Rank: 8, File: 2},
		{GameID: gameID, Color: Black, Name: "B", Rank: 8, File: 3},
		{GameID: gameID, Color: Black, Name: "Q", Rank: 8, File: 4},
		{GameID: gameID, Color: Black, Name: "K", Rank: 8, File: 5},
		{GameID: gameID, Color: Black, Name: "B", Rank: 8, File: 6},
		{GameID: gameID, Color: Black, Name: "N", Rank: 8, File: 7},
		{GameID: gameID, Color: Black, Name: "R", Rank: 8, File: 8},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 1},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 2},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 3},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 4},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 5},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 6},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 7},
		{GameID: gameID, Color: Black, Name: "P", Rank: 7, File: 8},
	}

	return pieces
}
