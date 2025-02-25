package models

import "time"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	ID        int64     `json:"game_id"`
	IsOver    bool      `json:"is_over"`
	CreatedAt time.Time `json:"created_at"`
}

// A struct for matching a User with a Game.
type Player struct {
	UserID int64 `json:"user_id"`
	GameID int64 `json:"game_id"`
	Color  bool  `json:"color"`
}

// A move. If viewing an old game, the Move has an ID. Otherwise it doesn't.
type Move struct {
	ID        int64     `json:"id"`
	GameID    int64     `json:"game_id"`
	Notation  string    `json:"notation"`
	CreatedAt time.Time `json:"created_at"`
}

type Piece struct {
	ID     int64  `json:"id"`
	GameID int64  `json:"game_id"`
	Color  bool   `json:"color"`
	Name   string `json:"name"`
	Rank   int    `json:"rank"`
	File   int    `json:"file"`
}

type GameOver struct {
	Draw        bool `json:"draw"`
	Checkmate   bool `json:"checkmate"`
	WinnerColor bool `json:"winner_color"`
}

type ValidationResult struct {
	IsValidMove bool     `json:"is_valid_move"`
	GameOver    GameOver `json:"game_situation"`
	KingInCheck bool     `json:"king_in_check"`
}

type PieceUpdate struct {
	DeletePiece bool  `json:"delete_piece"`
	Piece       Piece `json:"piece"`
}

// All the required data for a database saving entry.
type ChessEntry struct {
	GameID         int64         `json:"game_id"`
	Move           Move          `json:"move"`
	GameOver       GameOver      `json:"status"`
	AffectedPieces []PieceUpdate `json:"affected_pieces"`
}
