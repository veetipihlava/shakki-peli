package sqlite

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreateGame(whitePlayer int64, blackPlayer int64) error {
	query := `INSERT INTO games (white_player_id, black_player_id)
              VALUES (?, ?);`
	_, err := db.Connection.Exec(query, whitePlayer, blackPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ReadGame(firstPlayer int64, secondPlayer int64) (*models.Game, error) {
	query := `SELECT *
              FROM games 
			  WHERE (white_player_id = ? AND black_player_id = ?)
			   	 OR (white_player_id = ? AND black_player_id = ?);`
	row := db.Connection.QueryRow(query, firstPlayer, secondPlayer, secondPlayer, firstPlayer)

	game := models.Game{}
	err := row.Scan(&game.ID, &game.WhitePlayerID, &game.BlackPlayerID, &game.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &game, nil
}
