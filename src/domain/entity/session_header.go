package entity

type SessionHeader struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
