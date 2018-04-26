package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type StudentAnswer struct {
	ID              int64          `db:"id" json:"id"`
	Rational        string         `db:"rational" json:"rational"`
	QuestionID      int64          `db:"question_id" json:"question_id"`
	UserID          int64          `db:"user_id" json:"user_id"`
	RatingID        int64          `db:"rating_id" json:"rating_id"`
	InitialAnswerID int64          `db:"initial_answer_id" json:"initial_answer_id"`
	FinalAnswerID   int64          `db:"final_answer_id" json:"final_answer_id"`
	CreatedAt       time.Time      `db:"created_at" json:"-"`
	UpdatedAt       time.Time      `db:"updated_at" json:"-"`
	DeletedAt       mysql.NullTime `db:"deleted_at" json:"-"`
}
