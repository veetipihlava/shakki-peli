package sqlite

import "github.com/veetipihlava/shakki-peli/internal/models"

func (db *Database) CreatePiece(gameID int64, color bool, name string, rank int, file int) error {
	query := `INSERT INTO pieces (game_id, color, name, rank, file)
              VALUES (?, ?, ?, ?, ?);`
	_, err := db.Connection.Exec(query, gameID, color, name, rank, file)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ReadPieces(gameID int64) ([]models.Piece, error) {
	query := `SELECT *
              FROM pieces WHERE game_id = ?;`
	rows, err := db.Connection.Query(query, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []models.Piece
	for rows.Next() {
		piece := models.Piece{}
		err := rows.Scan(&piece.ID, &piece.GameID, &piece.Color, &piece.Name, &piece.Rank, &piece.File)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

func (db *Database) UpdatePiece(pieceID int64, rank int, file int) error {
	query := `UPDATE pieces 
			  SET rank = ?, file = ?
			  WHERE id = ?;`
	_, err := db.Connection.Exec(query, rank, file, pieceID)

	return err
}

func (db *Database) DeletePiece(pieceID int64) error {
	query := `DELETE FROM pieces 
			  WHERE id = ?;`
	_, err := db.Connection.Exec(query, pieceID)

	return err
}
