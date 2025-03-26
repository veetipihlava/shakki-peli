package redis

import (
	"encoding/json"
	"fmt"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (r *RedisClient) PublishEntry(chessEntry models.ChessEntry) error {
	channel := fmt.Sprintf("game:%d:entries", chessEntry.GameID)
	messageData, err := json.Marshal(chessEntry)
	if err != nil {
		return fmt.Errorf("failed to marshal entry message: %v", err)
	}

	err = r.Client.Publish(r.Ctx, channel, messageData).Err()
	if err != nil {
		return fmt.Errorf("failed to publish entry message: %v", err)
	}

	return nil
}
