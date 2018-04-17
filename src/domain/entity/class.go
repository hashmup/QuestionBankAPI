package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Class struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	SchoolID  int64          `db:"school_id" json:"school_id"`
	Term      string         `db:"term" json:"term"`
	ClassCode string         `db:"class_code" json:"class_code"`
	CreatedAt time.Time      `db:"created_at" json:"-"`
	UpdatedAt time.Time      `db:"updated_at" json:"-"`
	DeletedAt mysql.NullTime `db:"deleted_at" json:"-"`
}
