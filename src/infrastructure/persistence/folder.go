package persistence

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type folderRepository struct {
	DBClient *sqlx.DB
}

func NewFolderRepository(dbc *sqlx.DB) repository.FolderRepository {
	return &folderRepository{
		DBClient: dbc,
	}
}

func (repo *folderRepository) GetQuestions(ctx context.Context, userID, folderID int64) ([]*entity.QuestionRequest, error) {
	sql, args, _ := sq.Select("*").From("questions").Where(sq.Eq{"folder_id": folderID}).ToSql()
	questions := []entity.Question{}
	err := repo.DBClient.SelectContext(ctx, &questions, sql, args...)
	if err != nil {
		return nil, err
	}
	questionRequests := []*entity.QuestionRequest{}
	for i := range questions {
		answers := []entity.Answer{}
		sql, args, _ = sq.Select("*").From("answers").Where(sq.Eq{"id": []int64{questions[i].Answer1, questions[i].Answer2, questions[i].Answer3, questions[i].Answer4}}).ToSql()
		err := repo.DBClient.SelectContext(ctx, &answers, sql, args...)
		if err != nil {
			return nil, err
		}
		_answers := []*entity.Answer{}
		for j := range answers {
			_answers = append(_answers, &answers[j])
		}
		studentAnswer := entity.StudentAnswer{}
		sql, args, _ = sq.Select("*").From("student_answers").Where(sq.Eq{"user_id": userID, "question_id": questions[i].ID}).ToSql()
		err = repo.DBClient.SelectContext(ctx, &studentAnswer, sql, args...)

		questionRequests = append(questionRequests, &entity.QuestionRequest{
			Question:        questions[i].Question,
			Solved:          err != nil,
			Answers:         _answers,
			CorrectAnswerID: questions[i].CorrectAnswerID,
		})
	}
	return questionRequests, nil
}

func (repo *folderRepository) PostQuestions(ctx context.Context, profID, folderID int64, questionRequest *entity.QuestionRequest) error {
	now := time.Now()
	sql := "INSERT INTO answers (name, created_at, updated_at) VALUES(:name, :created_at, :updated_at)"
	answerIDs := []int64{}
	for i := range questionRequest.Answers {
		res, err := repo.DBClient.NamedExecContext(ctx, sql, questionRequest.Answers[i])
		if err != nil {
			return err
		}
		answerID, err := res.LastInsertId()
		answerIDs = append(answerIDs, answerID)
	}
	question := entity.Question{
		Question:        questionRequest.Question,
		FolderID:        folderID,
		ProfID:          profID,
		Answer1:         answerIDs[0],
		Answer2:         answerIDs[1],
		Answer3:         answerIDs[2],
		Answer4:         answerIDs[3],
		CorrectAnswerID: questionRequest.CorrectAnswerID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	sql = "INSERT INTO questions (question, answer_1, answer_2, answer_3, answer_4, correct_answer_id, created_at, updated_at) VALUES (:question, :answer_1, :answer_2, :answer_3, :answer_4, :correct_answer_id, :created_at, :updated_at)"
	_, err := repo.DBClient.NamedExecContext(ctx, sql, question)
	if err != nil {
		return err
	}
	return nil
}
