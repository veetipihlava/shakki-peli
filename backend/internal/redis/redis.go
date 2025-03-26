package redis

import (
	"github.com/veetipihlava/shakki-peli/internal/models"
)

type Redis interface {
	// Redis hash
	SaveGame(game *models.Game) error
	ReadGame(gameID int64) (*models.Game, error)
	SavePlayer(player *models.Player) error
	ReadPlayer(playerID int64, gameID int64) (*models.Player, error)
	SavePieces(pieces []models.Piece) error
	ReadPieces(gameID int64) ([]models.Piece, error)
	UpdatePiece(piece *models.Piece) error
	DeletePiece(piece *models.Piece) error

	// Redis stream
	SaveMove(move *models.Move) error
	GetMoves(gameID int64) ([]models.Move, error)

	// Redis pub/sub
	PublishEntry(chessEntry models.ChessEntry) error
}

type RedisService struct {
	redis Redis
}

func (rs *RedisService) SaveGame(game *models.Game) error {
	return rs.redis.SaveGame(game)
}

func (rs *RedisService) ReadGame(gameID int64) (*models.Game, error) {
	return rs.redis.ReadGame(gameID)
}

func (rs *RedisService) SavePlayer(player *models.Player) error {
	return rs.redis.SavePlayer(player)
}

func (rs *RedisService) ReadPlayer(playerID int64, gameID int64) (*models.Player, error) {
	return rs.redis.ReadPlayer(playerID, gameID)
}

func (rs *RedisService) SavePieces(pieces []models.Piece) error {
	return rs.redis.SavePieces(pieces)
}

func (rs *RedisService) ReadPieces(gameID int64) ([]models.Piece, error) {
	return rs.redis.ReadPieces(gameID)
}

func (rs *RedisService) UpdatePiece(piece *models.Piece) error {
	return rs.redis.UpdatePiece(piece)
}

func (rs *RedisService) DeletePiece(piece *models.Piece) error {
	return rs.redis.DeletePiece(piece)
}
