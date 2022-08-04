package models

import "time"

type Thread struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Subject   string    `json:"subject"`
	UserID    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
