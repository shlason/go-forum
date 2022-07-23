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

func (s *Session) UpdateByUserID() error {
	_, err := db.Exec("UPDATE sessions SET uuid = ?, expiry = ? WHERE id = ?", s.UUID, s.Expiry, s.ID)
	return err
}

func (s *Session) UpdateByUUID() error {
	_, err := db.Exec("UPDATE sessions SET uuid = ?, expiry = ? WHERE uuid = ?", s.UUID, s.Expiry, s.UUID)
	return err
}
