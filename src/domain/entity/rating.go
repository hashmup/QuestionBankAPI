package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Rating struct {
	ID        int64          `db:"id" json:"id"`
	Rating1   int            `db:"rating_1" json:"rating_1"`
	Rating2   int            `db:"rating_2" json:"rating_2"`
	Rating3   int            `db:"rating_3" json:"rating_3"`
	Rating4   int            `db:"rating_4" json:"rating_4"`
	CreatedAt time.Time      `db:"created_at" json:"-"`
	UpdatedAt time.Time      `db:"updated_at" json:"-"`
	DeletedAt mysql.NullTime `db:"deleted_at" json:"-"`
}
