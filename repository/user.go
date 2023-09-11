package repository

import (
	"database/sql"
	"time"

	"github.com/halosatrio/bebop/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) Store(user models.User) error {
	currentTime := time.Now()
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	query := `INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT email, password FROM users WHERE email=$1`
	err := r.DB.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
