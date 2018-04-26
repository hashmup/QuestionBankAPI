package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type QuestionRequest struct {
	Question        string    `json:"question"`
	Solved          bool      `json:"solved"`
	Answers         []*Answer `json:"answers"`
	CorrectAnswerID int64     `json:"correct_answer_id"`
}

type Question struct {
	ID              int64          `db:"id"`
	Question        string         `db:"question"`
	FolderID        int64          `db:"folder_id"`
	ProfID          int64          `db:"prof_id"`
	Answer1         int64          `db:"answer_1"`
	Answer2         int64          `db:"answer_2"`
	Answer3         int64          `db:"answer_3"`
	Answer4         int64          `db:"answer_4"`
	CorrectAnswerID int64          `db:"correct_answer_id"`
	CreatedAt       time.Time      `db:"created_at" json:"-"`
	UpdatedAt       time.Time      `db:"updated_at" json:"-"`
	DeletedAt       mysql.NullTime `db:"deleted_at" json:"-"`
}
