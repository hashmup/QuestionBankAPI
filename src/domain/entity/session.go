package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Session struct {
	ID        int64          `db:"id" json:"id"`
	UserID    int64          `db:"user_id" json:"user_id"`
	Token     string         `db:"token" json:"token"`
	ExpiredAt time.Time      `db:"expired_at" json:"expired_at"`
	CreatedAt time.Time      `db:"created_at" json:"-"`
	UpdatedAt time.Time      `db:"updated_at" json:"-"`
	DeletedAt mysql.NullTime `db:"deleted_at" json:"-"`
}
