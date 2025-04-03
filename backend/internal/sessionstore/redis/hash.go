package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

// Saves a game to games
func (r *Redis) SaveGame(game *models.Game) error {
	key := "games"
	field := strconv.FormatInt(game.ID, 10)
	value, err := json.Marshal(game)
	if err != nil {
		return err
	}

	err = r.Client.HSet(r.Ctx, key, field, value).Err()
	if err != nil {
		return err
	}

	return nil
}

// Reads a game from games and returns it
func (r *Redis) ReadGame(gameID int64) (*models.Game, error) {
	key := "games"
	field := strconv.FormatInt(gameID, 10)

	gameData, err := r.Client.HGet(r.Ctx, key, field).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("game with ID %d not found", gameID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to read game: %v", err)
	}

	var game models.Game
	err = json.Unmarshal([]byte(gameData), &game)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal game data: %v", err)
	}

	return &game, nil
}

// Removes a game from games, and ensures all data related to this game is removed
func (r *Redis) RemoveGame(gameID int64) error {
	// 1. Remove this game from games
	gameKey := "games"
	gameField := strconv.FormatInt(gameID, 10)
	if err := r.Client.HDel(r.Ctx, gameKey, gameField).Err(); err != nil {
		return fmt.Errorf("failed to remove game with ID %d: %v", gameID, err)
	}

	// 2. Remove the remaining pieces from games:<game_id>:pieces
	pieceKey := fmt.Sprintf("games:%d:pieces", gameID)
	pieceFields, err := r.Client.HKeys(r.Ctx, pieceKey).Result()
	if err != nil {
		log.Printf("[SESSION STORE] Could not fetch pieces for game %d: %v", gameID, err)
	} else {
		if len(pieceFields) > 0 {
			if err := r.Client.HDel(r.Ctx, pieceKey, pieceFields...).Err(); err != nil {
				log.Printf("[SESSION STORE] Failed to remove pieces from game %d: %v", gameID, err)
			}
		}
		r.Client.Del(r.Ctx, pieceKey)
	}

	// 3. Ensure all players are removed from games:<game_id>:players
	playerKey := fmt.Sprintf("games:%d:players", gameID)
	playerFields, err := r.Client.HKeys(r.Ctx, playerKey).Result()
	if err != nil {
		log.Printf("[SESSION STORE] Could not fetch players for game %d: %v", gameID, err)
	} else {
		if len(playerFields) > 0 {
			if err := r.Client.HDel(r.Ctx, playerKey, playerFields...).Err(); err != nil {
				log.Printf("[SESSION STORE] Failed to remove players from game %d: %v", gameID, err)
			}
		}
		r.Client.Del(r.Ctx, playerKey)
	}

	log.Printf("[SESSION STORE] Game %d fully removed (game, players, pieces)", gameID)
	return nil
}

// Saves a player to games:game_id:players
func (r *Redis) SavePlayer(player *models.Player) error {
	key := fmt.Sprintf("games:%d:players", player.GameID)
	field := strconv.FormatInt(player.UserID, 10)
	value, err := json.Marshal(player)

	if err != nil {
		return fmt.Errorf("failed to marshal player: %v", err)
	}

	err = r.Client.HSet(r.Ctx, key, field, value).Err()
	if err != nil {
		return fmt.Errorf("failed to save player: %v", err)
	}

	return nil
}

// Reads a single player from games:game_id:players and returns it
func (r *Redis) ReadPlayer(playerID int64, gameID int64) (*models.Player, error) {
	key := fmt.Sprintf("games:%d:players", gameID)
	field := strconv.FormatInt(playerID, 10)

	playerData, err := r.Client.HGet(r.Ctx, key, field).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("player with ID %d in game %d not found", playerID, gameID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to read player: %v", err)
	}

	var player models.Player
	err = json.Unmarshal([]byte(playerData), &player)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal player data: %v", err)
	}

	return &player, nil
}

// Reads all players from games:game_id:players and returns them as a slice
func (r *Redis) ReadPlayers(gameID int64) ([]models.Player, error) {
	key := fmt.Sprintf("games:%d:players", gameID)

	playersData, err := r.Client.HGetAll(r.Ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to read players for game %d: %v", gameID, err)
	}

	var players []models.Player
	for _, value := range playersData {
		var player models.Player
		err := json.Unmarshal([]byte(value), &player)
		if err != nil {
			log.Printf("failed to unmarshal player data in game %d: %v", gameID, err)
			continue
		}
		players = append(players, player)
	}

	return players, nil
}

func (r *Redis) RemovePlayer(playerID int64, gameID int64) error {
	key := fmt.Sprintf("games:%d:players", gameID)
	field := strconv.FormatInt(playerID, 10)

	if err := r.Client.HDel(r.Ctx, key, field).Err(); err != nil {
		return fmt.Errorf("failed to remove player with ID %d from game %d: %v", playerID, gameID, err)
	}
	return nil
}

// Adds multiple pieces to games:game_id:pieces
func (r *Redis) SavePieces(pieces []models.Piece) error {
	if len(pieces) == 0 {
		return nil
	}

	key := fmt.Sprintf("games:%d:pieces", pieces[0].GameID)

	for _, piece := range pieces {
		field := strconv.FormatInt(piece.ID, 10)
		value, err := json.Marshal(piece)
		if err != nil {
			return fmt.Errorf("failed to marshal piece with ID %d: %v", piece.ID, err)
		}

		err = r.Client.HSet(r.Ctx, key, field, value).Err()
		if err != nil {
			return fmt.Errorf("failed to save piece with ID %d: %v", piece.ID, err)
		}
	}

	return nil
}

// Reads and returns a list of pieces from games:game_id:pieces
func (r *Redis) ReadPieces(gameID int64) ([]models.Piece, error) {
	key := fmt.Sprintf("games:%d:pieces", gameID)

	piecesData, err := r.Client.HGetAll(r.Ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to read pieces: %v", err)
	}

	var pieces []models.Piece
	for _, value := range piecesData {
		var piece models.Piece
		err := json.Unmarshal([]byte(value), &piece)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal piece: %v", err)
		}
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

// Overwrites the provided piece in games:game_id:pieces
func (r *Redis) UpdatePiece(piece *models.Piece) (*models.Piece, error) {
	key := fmt.Sprintf("games:%d:pieces", piece.GameID)
	field := strconv.FormatInt(piece.ID, 10)

	value, err := json.Marshal(piece)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal piece: %v", err)
	}

	if err := r.Client.HSet(r.Ctx, key, field, value).Err(); err != nil {
		return nil, fmt.Errorf("failed to update piece: %v", err)
	}

	return piece, nil
}

// Deletes a piece from games:game_id:pieces
func (r *Redis) RemovePiece(piece *models.Piece) error {
	key := fmt.Sprintf("games:%d:pieces", piece.GameID)
	field := strconv.FormatInt(piece.ID, 10)

	err := r.Client.HDel(r.Ctx, key, field).Err()
	if err != nil {
		return fmt.Errorf("failed to delete piece with ID %d: %v", piece.ID, err)
	}

	return nil
}
