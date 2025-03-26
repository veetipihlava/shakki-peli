package updater

import (
	"encoding/json"
	"fmt"

	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/models"
	"github.com/veetipihlava/shakki-peli/internal/redis"
)

// Subscribe to a game specific channel and start listening to game entries
func ProcessGameEntries(r *redis.RedisClient, gameID int64) error {
	channel := fmt.Sprintf("game:%d:entries", gameID)
	pubsub := r.Client.Subscribe(r.Ctx, channel)

	_, err := pubsub.Receive(r.Ctx)
	if err != nil {
		return fmt.Errorf("failed to subscribe to channel: %v", err)
	}

	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var entry models.ChessEntry
			err := json.Unmarshal([]byte(msg.Payload), &entry)
			if err != nil {
				fmt.Printf("Failed to unmarshal message: %v", err)
				continue
			}
		}
	}()

	return nil
}

func WriteGameEntry(db *database.DatabaseService, entry models.ChessEntry) error {
	// Save the move in the database
	err := db.CreateMove(entry.GameID, entry.Move.Notation)
	if err != nil {
		return fmt.Errorf("failed to create move: %v", err)
	}

	// Update or delete affected pieces
	for _, pieceUpdate := range entry.AffectedPieces {
		if pieceUpdate.DeletePiece {
			// Delete the piece from the database
			err := db.DeletePiece(pieceUpdate.Piece.ID)
			if err != nil {
				return fmt.Errorf("failed to delete piece with ID %d: %v", pieceUpdate.Piece.ID, err)
			}
		} else {
			// Update the piece in the database
			err := db.UpdatePiece(pieceUpdate.Piece)
			if err != nil {
				return fmt.Errorf("failed to update piece with ID %d: %v", pieceUpdate.Piece.ID, err)
			}
		}
	}

	// If the game is over (checkmate or draw), update the game status
	if entry.GameOver.Checkmate || entry.GameOver.Draw {
		game, err := db.ReadGame(entry.GameID)
		if err != nil {
			return fmt.Errorf("failed to read game for update: %v", err)
		}

		game.IsOver = true
		err = db.Database.EndGame(entry.GameID)
		if err != nil {
			return fmt.Errorf("failed to update game status: %v", err)
		}
	}

	return nil
}
