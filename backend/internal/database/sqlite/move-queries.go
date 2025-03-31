package sqlite

import "github.com/veetipihlava/shakki-peli/internal/models"

func (db *Database) CreateMove(gameID int64, notation string) (*models.Move, error) {
	query := `INSERT INTO moves (game_id, notation) VALUES (?, ?);`
	result, err := db.Connection.Exec(query, gameID, notation)
	if err != nil {
		return nil, err
	}

	moveID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return db.ReadMove(moveID)
}

func (db *Database) ReadMove(moveID int64) (*models.Move, error) {
	query := `SELECT * FROM moves WHERE id = ?;`
	row := db.Connection.QueryRow(query, moveID)

	var move models.Move
	err := row.Scan(&move.ID, &move.GameID, &move.Notation, &move.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &move, nil
}

func (db *Database) ReadMoves(gameID int64) ([]models.Move, error) {
	query := `SELECT * 
              FROM moves WHERE game_id = ?;`
	rows, err := db.Connection.Query(query, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []models.Move
	for rows.Next() {
		move := models.Move{}
		err := rows.Scan(&move.ID, &move.GameID, &move.Notation, &move.CreatedAt)
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}

	return moves, nil
}
