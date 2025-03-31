package sqlite

import (
	"errors"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreatePieces(gameID int64, pieces []models.Piece) ([]models.Piece, error) {
	if len(pieces) == 0 {
		return nil, errors.New("no pieces in input")
	}

	query := `INSERT INTO pieces (game_id, color, name, rank, file) VALUES `
	vals := []interface{}{}

	for _, piece := range pieces {
		query += "(?, ?, ?, ?, ?),"
		vals = append(vals, piece.GameID, piece.Color, piece.Name, piece.Rank, piece.File)
	}
	query = query[:len(query)-1]

	_, err := db.Connection.Exec(query, vals...)
	if err != nil {
		return nil, err
	}

	return db.ReadPieces(gameID)
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
		err := rows.Scan(
			&piece.ID,
			&piece.GameID,
			&piece.Color,
			&piece.Name,
			&piece.Rank,
			&piece.File,
		)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

func (db *Database) UpdatePiece(piece models.Piece) error {
	query := `UPDATE pieces 
			  SET game_id = ?, color = ?, name = ?, rank = ?, file = ?
			  WHERE id = ?;`
	_, err := db.Connection.Exec(query, piece.GameID, piece.Color, piece.Name, piece.Rank, piece.File, piece.ID)

	return err
}

func (db *Database) DeletePiece(pieceID int64) error {
	query := `DELETE FROM pieces 
			  WHERE id = ?;`
	_, err := db.Connection.Exec(query, pieceID)

	return err
}
