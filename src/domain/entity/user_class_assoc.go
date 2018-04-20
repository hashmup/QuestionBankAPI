package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type UserClassAssoc struct {
	ID        int64          `db:"id" json:"id"`
	UserID    int64          `db:"user_id" json:"user_id"`
	ClassID   int64          `db:"class_id" json:"class_id"`
	CreatedAt time.Time      `db:"created_at" json:"-"`
	UpdatedAt time.Time      `db:"updated_at" json:"-"`
	DeletedAt mysql.NullTime `db:"deleted_at" json:"-"`
}
