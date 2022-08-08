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

func (p *Post) ReadAll() ([]Post, error) {
	rows, err := db.Query("SELECT id, uuid, content, user_id, thread_id, created_at, updated_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.UUID, &post.Content, &post.UserID, &post.ThreadID, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *Post) ReadAllByThreadID() ([]Post, error) {
	rows, err := db.Query("SELECT id, uuid, content, user_id, thread_id, created_at, updated_at FROM posts WHERE thread_id = ?", p.ThreadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.UUID, &post.Content, &post.UserID, &post.ThreadID, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
