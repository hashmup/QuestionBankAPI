package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	CreatedAt time.Time      `db:"created_at" json:"-"`
	UpdatedAt time.Time      `db:"updated_at" json:"-"`
	DeletedAt mysql.NullTime `db:"deleted_at" json:"-"`
}
