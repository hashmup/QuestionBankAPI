package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type QuestionRequest struct {
	QuestionID      int64     `json:"question_id"`
	Question        string    `json:"question"`
	Solved          bool      `json:"solved"`
	Answers         []*Answer `json:"answers"`
	Tags            []string  `json:"tags"`
	CorrectAnswerID int64     `json:"correct_answer_id"`
}

type Question struct {
	ID              int64          `db:"id" json:"id"`
	Question        string         `db:"question" json:"question"`
	FolderID        int64          `db:"folder_id" json:"folder_id"`
	InstructorID    int64          `db:"instructor_id" json:"instructor_id"`
	Answer1         int64          `db:"answer_1" json:"answer_1"`
	Answer2         int64          `db:"answer_2" json:"answer_2"`
	Answer3         int64          `db:"answer_3" json:"answer_3"`
	Answer4         int64          `db:"answer_4" json:"answer_4"`
	CorrectAnswerID int64          `db:"correct_answer_id" json:"correct_answer_id"`
	CreatedAt       time.Time      `db:"created_at" json:"-"`
	UpdatedAt       time.Time      `db:"updated_at" json:"-"`
	DeletedAt       mysql.NullTime `db:"deleted_at" json:"-"`
}

type QuestionAnswer struct {
	QuestionID      int64   `json:"question_id"`
	Rationale       string  `json:"rationale"`
	InitialAnswerID int64   `json:"initial_answer_id"`
	FinalAnswerID   int64   `json:"final_answer_id"`
	Rating          *Rating `json:"rating"`
}
