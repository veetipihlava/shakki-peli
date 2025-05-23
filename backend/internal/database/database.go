package database

import "github.com/veetipihlava/shakki-peli/internal/models"

// Interface for database interactions.
type Database interface {
	CreateUser(name string) (*models.User, error)
	ReadUser(userID int64) (*models.User, error)

	CreatePlayer(userID int64, gameID int64) (*models.Player, error)
	ReadPlayer(userID int64, gameID int64) (*models.Player, error)

	CreateGame() (*models.Game, error)
	ReadGame(gameID int64) (*models.Game, error)
	EndGame(gameID int64) error
	GetFullGameState(gameID int64) (*models.Game, []models.Player, []models.Piece, []models.Move, error)

	CreateMove(gameID int64, notation string) (*models.Move, error)
	ReadMoves(gameID int64) ([]models.Move, error)

	CreatePieces(gameID int64, pieces []models.Piece) ([]models.Piece, error)
	ReadPieces(gameID int64) ([]models.Piece, error)
	UpdatePiece(piece models.Piece) error
	DeletePiece(pieceID int64) error
}

// The database service for providing access to databases.
type DatabaseService struct {
	Database Database
}

// Creates a new user.
func (db *DatabaseService) CreateUser(name string) (*models.User, error) {
	return db.Database.CreateUser(name)
}

// Reads a user.
func (db *DatabaseService) ReadUser(userID int64) (*models.User, error) {
	return db.Database.ReadUser(userID)
}

// Creates a new player.
func (db *DatabaseService) CreatePlayer(userID int64, gameID int64) (*models.Player, error) {
	return db.Database.CreatePlayer(userID, gameID)
}

// Reads a player.
func (db *DatabaseService) ReadPlayer(userID int64, gameID int64) (*models.Player, error) {
	return db.Database.ReadPlayer(userID, gameID)
}

// Creates a new game.
func (db *DatabaseService) CreateGame() (*models.Game, error) {
	return db.Database.CreateGame()
}

// Reads a game.
func (db *DatabaseService) ReadGame(gameID int64) (*models.Game, error) {
	return db.Database.ReadGame(gameID)
}

// Changes the status of a game.
func (db *DatabaseService) EndGame(gameID int64) error {
	return db.Database.EndGame(gameID)
}

// Return everything related to this game
func (db *DatabaseService) GetFullGameState(gameID int64) (*models.Game, []models.Player, []models.Piece, []models.Move, error) {
	return db.Database.GetFullGameState(gameID)
}

// Creates a new move.
func (db *DatabaseService) CreateMove(gameID int64, notation string) (*models.Move, error) {
	return db.Database.CreateMove(gameID, notation)
}

// Reads a move.
func (db *DatabaseService) ReadMoves(gameID int64) ([]models.Move, error) {
	return db.Database.ReadMoves(gameID)
}

// Creates a new piece.
func (db *DatabaseService) CreatePieces(gameID int64, pieces []models.Piece) ([]models.Piece, error) {
	return db.Database.CreatePieces(gameID, pieces)
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
