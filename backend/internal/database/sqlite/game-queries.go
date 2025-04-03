package sqlite

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreateGame() (*models.Game, error) {
	query := `INSERT INTO games (is_over) VALUES (0);`
	result, err := db.Connection.Exec(query)
	if err != nil {
		return nil, err
	}

	gameID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return db.ReadGame(gameID)
}

func (db *Database) ReadGame(gameID int64) (*models.Game, error) {
	query := `SELECT *
              FROM games 
              WHERE id = ?;`
	row := db.Connection.QueryRow(query, gameID)

	game := models.Game{}
	err := row.Scan(&game.ID, &game.IsOver, &game.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &game, nil
}

func (db *Database) EndGame(gameID int64) error {
	query := `UPDATE games
              SET is_over = 1
              WHERE id = ?;`
	result, err := db.Connection.Exec(query, gameID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (db *Database) GetFullGameState(gameID int64) (*models.Game, []models.Player, []models.Piece, []models.Move, error) {
	game, err := db.ReadGame(gameID)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if game == nil {
		return nil, nil, nil, nil, nil
	}

	players, err := db.GetGamePlayers(gameID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	pieces, err := db.ReadPieces(gameID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	moves, err := db.ReadMoves(gameID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return game, players, pieces, moves, nil
}
