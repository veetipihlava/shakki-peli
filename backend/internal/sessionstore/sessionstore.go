package sessionstore

import "github.com/veetipihlava/shakki-peli/internal/models"

type SessionStore interface {
	// A SessionStore needs to:

	// Save, read and remove a game
	SaveGame(game *models.Game) error
	ReadGame(gameID int64) (*models.Game, error)
	RemoveGame(gameID int64) error

	// Save, read and remove a player
	SavePlayer(player *models.Player) error
	ReadPlayer(playerID int64, gameID int64) (*models.Player, error)
	RemovePlayer(playerID int64, gameID int64) error
	ReadPlayers(gameID int64) ([]models.Player, error)

	// Save, read, update and remove pieces
	SavePieces(pieces []models.Piece) error
	ReadPieces(gameID int64) ([]models.Piece, error)
	UpdatePiece(piece *models.Piece) (*models.Piece, error)
	RemovePiece(piece *models.Piece) error

	// Save, Get and Remove Moves for a game
	SaveMove(move models.Move) error
	GetMoves(gameID int64) ([]models.Move, error)
	RemoveMoves(gameID int64) error

	// Process a valid Entry
	PublishEntry(chessEntry models.ChessEntry) error
}

type SessionStoreService struct {
	SessionStore SessionStore
}

func (rs *SessionStoreService) SaveGame(game *models.Game) error {
	return rs.SessionStore.SaveGame(game)
}

func (rs *SessionStoreService) ReadGame(gameID int64) (*models.Game, error) {
	return rs.SessionStore.ReadGame(gameID)
}

func (rs *SessionStoreService) RemoveGame(gameID int64) error {
	return rs.SessionStore.RemoveGame(gameID)
}

func (rs *SessionStoreService) SavePlayer(player *models.Player) error {
	return rs.SessionStore.SavePlayer(player)
}

func (rs *SessionStoreService) ReadPlayer(playerID int64, gameID int64) (*models.Player, error) {
	return rs.SessionStore.ReadPlayer(playerID, gameID)
}

func (rs *SessionStoreService) ReadPlayers(gameID int64) ([]models.Player, error) {
	return rs.SessionStore.ReadPlayers(gameID)
}

func (rs *SessionStoreService) RemovePlayer(playerID int64, gameID int64) error {
	return rs.SessionStore.RemovePlayer(playerID, gameID)
}

func (rs *SessionStoreService) SavePieces(pieces []models.Piece) error {
	return rs.SessionStore.SavePieces(pieces)
}

func (rs *SessionStoreService) ReadPieces(gameID int64) ([]models.Piece, error) {
	return rs.SessionStore.ReadPieces(gameID)
}

func (rs *SessionStoreService) UpdatePiece(piece *models.Piece) (*models.Piece, error) {
	return rs.SessionStore.UpdatePiece(piece)
}

func (rs *SessionStoreService) RemovePiece(piece *models.Piece) error {
	return rs.SessionStore.RemovePiece(piece)
}

func (rs *SessionStoreService) SaveMove(move models.Move) error {
	return rs.SessionStore.SaveMove(move)
}

func (rs *SessionStoreService) GetMoves(gameID int64) ([]models.Move, error) {
	return rs.SessionStore.GetMoves(gameID)
}

func (rs *SessionStoreService) RemoveMoves(gameID int64) error {
	return rs.SessionStore.RemoveMoves(gameID)
}

func (rs *SessionStoreService) PublishEntry(chessEntry models.ChessEntry) error {
	return rs.SessionStore.PublishEntry(chessEntry)
}
