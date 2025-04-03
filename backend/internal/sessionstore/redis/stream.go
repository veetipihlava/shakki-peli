package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (r *Redis) SaveMove(move models.Move) error {
	key := fmt.Sprintf("games:%d:moves", move.GameID)

	_, err := r.Client.XAdd(r.Ctx, &redis.XAddArgs{
		Stream: key,
		Values: map[string]interface{}{
			"id":        move.ID,
			"notation":  move.Notation,
			"createdAt": move.CreatedAt.Unix(),
		},
	}).Result()

	if err != nil {
		return fmt.Errorf("failed to save move: %v", err)
	}
	return nil
}

func (r *Redis) GetMoves(gameID int64) ([]models.Move, error) {
	key := fmt.Sprintf("games:%d:moves", gameID)

	movesData, err := r.Client.XRange(r.Ctx, key, "-", "+").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get moves: %v", err)
	}

	var moves []models.Move
	for _, entry := range movesData {
		move := models.Move{
			GameID: gameID,
		}
		if id, ok := entry.Values["id"].(string); ok {
			move.ID, _ = strconv.ParseInt(id, 10, 64)
		}
		if notation, ok := entry.Values["notation"].(string); ok {
			move.Notation = notation
		}
		if createdAtStr, ok := entry.Values["createdAt"].(string); ok {
			timestamp, _ := strconv.ParseInt(createdAtStr, 10, 64)
			move.CreatedAt = time.Unix(timestamp, 0)
		}

		moves = append(moves, move)
	}

	return moves, nil
}

func (r *Redis) RemoveMoves(gameID int64) error {
	key := fmt.Sprintf("games:%d:moves", gameID)

	err := r.Client.Del(r.Ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to flush moves for game %d: %v", gameID, err)
	}

	return nil
}
