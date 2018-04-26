package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Rating struct {
	ID           int64          `db:"id" json:"id"`
	Confidence   int            `db:"confidence" json:"confidence"`
	Difficulty   int            `db:"difficulty" json:"diffuculty"`
	Clarity      int            `db:"clarity" json:"clarity"`
	Preparedness int            `db:"preparedness" json:"preparedness"`
	CreatedAt    time.Time      `db:"created_at" json:"-"`
	UpdatedAt    time.Time      `db:"updated_at" json:"-"`
	DeletedAt    mysql.NullTime `db:"deleted_at" json:"-"`
}
