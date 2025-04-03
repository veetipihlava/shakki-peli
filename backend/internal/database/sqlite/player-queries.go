package sqlite

import (
	"database/sql"
	"errors"
	"log"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreatePlayer(userID int64, gameID int64) (*models.Player, error) {
	query := `SELECT COUNT(*) FROM players WHERE game_id = ?;`
	var count int
	err := db.Connection.QueryRow(query, gameID).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count >= 2 {
		return nil, errors.New("game already has two players")
	}

	color := true
	if count == 1 {
		color = false
	}

	insertQuery := `INSERT INTO players (user_id, game_id, color) VALUES (?, ?, ?);`
	_, err = db.Connection.Exec(insertQuery, userID, gameID, color)
	if err != nil {
		return nil, err
	}

	return db.ReadPlayer(userID, gameID)
}

func (db *Database) ReadPlayer(userID int64, gameID int64) (*models.Player, error) {
	query := `SELECT * FROM players WHERE user_id = ? AND game_id = ? ;`

	row := db.Connection.QueryRow(query, userID, gameID)

	var player models.Player
	err := row.Scan(
		&player.UserID,
		&player.GameID,
		&player.Color,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	log.Println("[DB]: Player in game ", player.GameID)
	return &player, nil
}

func (db *Database) GetGamePlayers(gameID int64) ([]models.Player, error) {
	query := `SELECT * FROM players WHERE game_id = ?;`

	rows, err := db.Connection.Query(query, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.UserID, &player.GameID, &player.Color)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}
