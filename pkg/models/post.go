package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Content   string    `json:"content"`
	UserID    int       `json:"userId"`
	ThreadID  int       `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *Post) Create() error {
	_, err := db.Exec("INSERT INTO posts (uuid, content, user_id, thread_id) VALUES (?, ?, ?, ?)", uuid.New().String(), p.Content, p.UserID, p.ThreadID)
	return err
}
