package models

import (
	"time"

	"github.com/shlason/go-forum/pkg/utils"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAte"`
}

func (u *User) Create() error {
	hp, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, hp)

	return err
}

func (u *User) ReadAll() ([]User, error) {
	rows, err := db.Query("SELECT id, email, password, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) ReadByName() error {
	return db.QueryRow("SELECT id, email, password, created_at, updated_at FROM users WHERE name = ?", u.Name).
		Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT id, name, password, created_at, updated_at FROM users WHERE email = ?", u.Email).
		Scan(&u.ID, &u.Name, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ReadByUserID() error {
	return db.QueryRow("SELECT name, email, password, created_at, updated_at FROM users WHERE id = ?", u.ID).
		Scan(&u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}
