package models

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Subject   string    `json:"subject"`
	UserID    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *Thread) Create() error {
	_, err := db.Exec("INSERT INTO threads (uuid, subject, user_id) VALUES (?, ?, ?)", uuid.New().String(), t.Subject, t.UserID)
	return err
}

func (t *Thread) ReadAll() ([]Thread, error) {
	rows, err := db.Query("SELECT id, uuid, subject, user_id, created_at, updated_at FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []Thread
	for rows.Next() {
		thread := Thread{}
		if err := rows.Scan(&thread.ID, &thread.UUID, &thread.Subject, &thread.UserID, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func (t *Thread) ReadByID() error {
	return db.QueryRow("SELECT uuid, subject, user_id, created_at, updated_at FROM threads WHERE id = ?", t.ID).Scan(
		&t.UUID, &t.Subject, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
}
