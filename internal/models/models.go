package models

import "time"

type Player struct {
	ID    int64
	Name  string
	Color bool
}

type Game struct {
	ID            int64
	WhitePlayerID int64
	BlackPlayerID int64
	CreatedAt     time.Time
}

type Move struct {
	ID        int64
	GameID    int64
	Notation  string
	CreatedAt time.Time
}

type Piece struct {
	ID     int64
	GameID int64
	Color  bool
	Name   string
	Rank   int
	File   int
}
