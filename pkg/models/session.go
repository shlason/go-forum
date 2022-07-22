package models

type Session struct {
	ID     int
	UUID   string
	UserID int
}

func (s *Session) Create() error {
	_, err := db.Exec("INSERT ONTO sessions (uuid, user_id) VALUES (?, ?)", s.UUID, s.UserID)
	return err
}
