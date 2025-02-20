package sqlite

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreatePlayer(userID int64, gameID int64, color bool) error {
	query := `INSERT INTO players (game_id, user_id, color)
			  VALUES (?, ?, ?);`
	_, err := db.Connection.Exec(query, gameID, userID, color)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ReadPlayer(userID int64, gameID int64) (*models.Player, error) {
	query := `SELECT *
			  FROM players
			  WHERE 
			  	user_id = ? AND game_id = ? ;`

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

	return &player, nil
}
