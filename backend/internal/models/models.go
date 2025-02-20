package models

import "time"

type User struct {
	ID   int64
	Name string
}

type Game struct {
	ID        int64
	IsOver    bool
	CreatedAt time.Time
}

type Player struct {
	UserID int64
	GameID int64
	Color  bool
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
