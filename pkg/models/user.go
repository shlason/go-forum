package models

import (
	"time"

	"github.com/shlason/go-forum/pkg/utils"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Create() error {
	hp, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, hp)

	return err
}

func (u *User) ReadByName() error {
	return db.QueryRow("SELECT id, email, password, created_at, updated_at FROM users WHERE name = ?", u.Name).
		Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT id, name, password, created_at, updated_at FROM users WHERE email = ?", u.Email).
		Scan(&u.ID, &u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}
