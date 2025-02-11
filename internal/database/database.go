package database

import "github.com/veetipihlava/shakki-peli/internal/models"

// Interface for database interactions.
type Database interface {
	CreatePlayer(name string, color bool) (int64, error)
	ReadPlayer(playerID int64) (*models.Player, error)

	CreateGame(whitePlayer int64, blackPlayer int64) error
	ReadGame(firstPlayer int64, secondPlayer int64) (*models.Game, error)

	CreateMove(gameID int64, notation string) error
	ReadMoves(gameID int64) ([]models.Move, error)

	CreatePiece(gameID int64, color bool, name string, rank int, file int) error
	ReadPieces(gameID int64) ([]models.Piece, error)
	UpdatePiece(pieceID int64, rank int, file int) error
	DeletePiece(pieceID int64) error
}

// The database service for providing access to databases.
type DatabaseService struct {
	Database Database
}

// Creates a new player.
func (db *DatabaseService) CreatePlayer(name string, color bool) (int64, error) {
	return db.Database.CreatePlayer(name, color)
}

// Reads a player.
func (db *DatabaseService) ReadPlayer(playerID int64) (*models.Player, error) {
	return db.Database.ReadPlayer(playerID)
}

// Creates a new game.
func (db *DatabaseService) CreateGame(whitePlayer int64, blackPlayer int64) error {
	return db.Database.CreateGame(whitePlayer, blackPlayer)
}

// Reads a game.
func (db *DatabaseService) ReadGame(firstPlayer int64, secondPlayer int64) (*models.Game, error) {
	return db.Database.ReadGame(firstPlayer, secondPlayer)
}

// Creates a new move.
func (db *DatabaseService) CreateMove(gameID int64, notation string) error {
	return db.Database.CreateMove(gameID, notation)
}

// Reads a move.
func (db *DatabaseService) ReadMoves(gameID int64) ([]models.Move, error) {
	return db.Database.ReadMoves(gameID)
}

// Creates a new piece.
func (db *DatabaseService) CreatePiece(gameID int64, color bool, name string, rank int, file int) error {
	return db.Database.CreatePiece(gameID, color, name, rank, file)
}

// Reads a piece.
func (db *DatabaseService) ReadPieces(gameID int64) ([]models.Piece, error) {
	return db.Database.ReadPieces(gameID)
}

// Updates a piece.
func (db *DatabaseService) UpdatePiece(pieceID int64, rank int, file int) error {
	return db.Database.UpdatePiece(pieceID, rank, file)
}

// Deletes a piece.
func (db *DatabaseService) DeletePiece(pieceID int64) error {
	return db.Database.DeletePiece(pieceID)
}
