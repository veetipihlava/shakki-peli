package sqlite

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreatePlayer(name string, color bool) (int64, error) {
	query := `INSERT INTO players (name, color)
			  VALUES (?, ?);`
	result, err := db.Connection.Exec(query, name, color)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *Database) ReadPlayer(playerID int64) (*models.Player, error) {
	query := `SELECT *
			  FROM players WHERE id = ?;`
	row := db.Connection.QueryRow(query, playerID)

	player := models.Player{}
	err := row.Scan(&player.ID, &player.Name, &player.Color)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &player, nil
}
