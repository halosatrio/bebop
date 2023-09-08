package repository

import (
	"database/sql"

	"github.com/halosatrio/bebop/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	// SQL Insert with hashed password
	query := `INSERT INTO users(email, password) VALUES($1, $2)`
	_, err := db.Exec(query, user.Email, user.Password)
	return err
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
