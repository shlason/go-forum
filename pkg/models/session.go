package models

import "time"

type Session struct {
	ID     int
	UUID   string
	UserID int
	Expiry time.Time
}

func (s *Session) Create() error {
	_, err := db.Exec("INSERT ONTO sessions (uuid, user_id) VALUES (?, ?)", s.UUID, s.UserID)
	return err
}

func (s *Session) ReadByUserID() error {
	return db.QueryRow("SELECT id, uuid FROM sessions WHERE user_id = ?", s.UserID).
		Scan(&s.ID, &s.UUID)
}
