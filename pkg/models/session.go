package models

import "time"

type Session struct {
	ID     int
	UUID   string
	UserID int
	Expiry time.Time
}

func (s *Session) Create() error {
	_, err := db.Exec("INSERT INTO sessions (uuid, user_id, expiry) VALUES (?, ?, ?)", s.UUID, s.UserID, s.Expiry)
	return err
}

func (s *Session) ReadByUserID() error {
	return db.QueryRow("SELECT id, uuid, expiry FROM sessions WHERE user_id = ?", s.UserID).
		Scan(&s.ID, &s.UUID, &s.Expiry)
}

func (s *Session) Update() error {
	_, err := db.Exec("UPDATE session SET uuid = ?, expiry = ? WHERE id = ?", s.UUID, s.Expiry, s.ID)
	return err
}
