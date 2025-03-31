package sqlite

import (
	"database/sql"

	"github.com/veetipihlava/shakki-peli/internal/models"
)

func (db *Database) CreateUser(name string) (*models.User, error) {
	query := `INSERT INTO users (name) VALUES (?);`
	result, err := db.Connection.Exec(query, name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return db.ReadUser(id)
}

func (db *Database) ReadUser(userID int64) (*models.User, error) {
	query := `SELECT *
			  FROM users
			  WHERE id = ?;`

	row := db.Connection.QueryRow(query, userID)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
