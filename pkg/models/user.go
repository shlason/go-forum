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
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ReadByName() error {
	return db.QueryRow("SELECT email, password, created_at, updated_at WHERE name = ?", u.Name).
		Scan(&u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT name, password, created_at, updated_at WHERE email = ?", u.Email).
		Scan(&u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}
