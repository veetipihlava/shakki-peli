package database

import "github.com/veetipihlava/shakki-peli/internal/models"

// Interface for database interactions.
type Database interface {
	CreateUser(name string) (int64, error)
	ReadUser(userID int64) (*models.User, error)

	CreatePlayer(userID int64, gameID int64, color bool) (int64, error)
	ReadPlayer(userID int64, gameID int64) (*models.Player, error)

	CreateGame() (int64, error)
	ReadGame(gameID int64) (*models.Game, error)

	CreateMove(gameID int64, notation string) error
	ReadMoves(gameID int64) ([]models.Move, error)

	CreatePieces(pieces []models.Piece) error
	ReadPieces(gameID int64) ([]models.Piece, error)
	UpdatePiece(piece models.Piece) error
	DeletePiece(pieceID int64) error
}

// The database service for providing access to databases.
type DatabaseService struct {
	Database Database
}

// Creates a new user.
func (db *DatabaseService) CreateUser(name string) (int64, error) {
	return db.Database.CreateUser(name)
}

// Reads a user.
func (db *DatabaseService) ReadUser(userID int64) (*models.User, error) {
	return db.Database.ReadUser(userID)
}

// Creates a new player.
func (db *DatabaseService) CreatePlayer(userID int64, gameID int64, color bool) (int64, error) {
	return db.Database.CreatePlayer(gameID, userID, color)
}

// Reads a player.
func (db *DatabaseService) ReadPlayer(userID int64, gameID int64) (*models.Player, error) {
	return db.Database.ReadPlayer(userID, gameID)
}

// Creates a new game.
func (db *DatabaseService) CreateGame() (int64, error) {
	return db.Database.CreateGame()
}

// Reads a game.
func (db *DatabaseService) ReadGame(gameID int64) (*models.Game, error) {
	return db.Database.ReadGame(gameID)
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
func (db *DatabaseService) CreatePieces(pieces []models.Piece) error {
	return db.Database.CreatePieces(pieces)
}

// Reads a piece.
func (db *DatabaseService) ReadPieces(gameID int64) ([]models.Piece, error) {
	return db.Database.ReadPieces(gameID)
}

// Updates a piece.
func (db *DatabaseService) UpdatePiece(piece models.Piece) error {
	return db.Database.UpdatePiece(piece)
}

// Deletes a piece.
func (db *DatabaseService) DeletePiece(pieceID int64) error {
	return db.Database.DeletePiece(pieceID)
}
