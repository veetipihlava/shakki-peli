package chess

import (
	"log"
	"strings"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

// The boolean for white player.
const White bool = true

// The boolean for black player.
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

func LogGameState(pieces []models.Piece, moves []models.Move) {
	var board [8][8]string

	// Initialize board with "0"
	for r := range board {
		for f := range board[r] {
			board[r][f] = "0"
		}
	}

	// Place pieces
	for _, p := range pieces {
		rankIndex := 8 - p.Rank
		fileIndex := p.File - 1

		char := p.Name
		if !p.Color {
			char = strings.ToLower(char)
		}

		board[rankIndex][fileIndex] = char
	}

	log.Println("====== CURRENT BOARD STATE ======")
	for i, row := range board {
		rank := 8 - i
		log.Printf("%d | %s", rank, strings.Join(row[:], " "))
	}
	log.Println("   ------------------------")
	log.Println("    A B C D E F G H")

	// Log moves
	log.Println("====== MOVE HISTORY ======")
	if len(moves) == 0 {
		log.Println("No moves made yet.")
	} else {
		for i, m := range moves {
			log.Printf("%d. %s", i+1, m.Notation)
		}
	}
	log.Println("=================================")
}
