// models/user.go
package models

type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password"` // Don't output password in JSON
}
