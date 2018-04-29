package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type QuestionTagAssoc struct {
	ID         int64          `db:"id" json:"id"`
	QuestionID int64          `db:"question_id" json:"question_id"`
	TagID      int64          `db:"tag_id" json:"tag_id"`
	CreatedAt  time.Time      `db:"created_at" json:"-"`
	UpdatedAt  time.Time      `db:"updated_at" json:"-"`
	DeletedAt  mysql.NullTime `db:"deleted_at" json:"-"`
}
