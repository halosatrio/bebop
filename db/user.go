// db/user.go
package db

import (
	"database/sql"

	"github.com/halosatrio/bebop/models"
)

func StoreUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO bebop_user (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}
