package persistence

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type questionRepository struct {
	DBClient *sqlx.DB
}

func NewQuestionRepository(dbc *sqlx.DB) repository.QuestionRepository {
	return &questionRepository{
		DBClient: dbc,
	}
}

func (repo *questionRepository) GetQuestionRationales(ctx context.Context, questionID, classID int64) ([]*entity.QuestionRationale, error) {
	sql, args, _ := sq.Select("student_answers.*").From("student_answers").Join("questions ON student_answers.question_id = questions.id").Join("folders ON questions.folder_id = folders.id").Join("classes ON folders.class_id = classes.id").Where(sq.Eq{"questions.id": questionID, "classes.id": classID}).ToSql()
	studentAnswers := []entity.StudentAnswer{}
	err := repo.DBClient.SelectContext(ctx, &studentAnswers, sql, args...)
	if err != nil {
		return nil, err
	}
	questionRationales := []*entity.QuestionRationale{}
	for i := range studentAnswers {
		questionRationales = append(questionRationales, &entity.QuestionRationale{
			Rationale: studentAnswers[i].Rationale,
			AnswerID:  studentAnswers[i].InitialAnswerID,
		})
	}
	return questionRationales, nil
}

func (repo *questionRepository) GetQuestionAnswer(ctx context.Context, userID, questionID int64) (*entity.StudentAnswerResponse, error) {
	sql, args, _ := sq.Select("*").From("student_answers").Where(sq.Eq{"question_id": questionID, "user_id": userID}).ToSql()
	studentAnswer := entity.StudentAnswer{}
	rating := entity.Rating{}
	err := repo.DBClient.GetContext(ctx, &studentAnswer, sql, args...)
	if err != nil {
		return nil, err
	}
	sql, args, _ = sq.Select("*").From("ratings").Where(sq.Eq{"id": studentAnswer.RatingID}).ToSql()
	err = repo.DBClient.GetContext(ctx, &rating, sql, args...)
	if err != nil {
		return nil, err
	}

	return &entity.StudentAnswerResponse{
		Rationale:       studentAnswer.Rationale,
		QuestionID:      studentAnswer.QuestionID,
		UserID:          studentAnswer.UserID,
		Rating:          &rating,
		InitialAnswerID: studentAnswer.InitialAnswerID,
		FinalAnswerID:   studentAnswer.FinalAnswerID,
	}, nil
}

func (repo *questionRepository) PostQuestions(ctx context.Context, userID int64, questionAnswer *entity.QuestionAnswer) error {
	now := time.Now()
	sql := "INSERT INTO ratings (rating_1, rating_2, rating_3, rating_4, created_at, updated_at) VALUES(:rating_1, :rating_2, :rating_3, :rating_4, :created_at, :updated_at)"
	questionAnswer.Rating.CreatedAt = now
	questionAnswer.Rating.UpdatedAt = now

	res, err := repo.DBClient.NamedExecContext(ctx, sql, questionAnswer.Rating)
	if err != nil {
		return err
	}
	ratingID, err := res.LastInsertId()

	studentAnswer := entity.StudentAnswer{
		Rationale:       questionAnswer.Rationale,
		QuestionID:      questionAnswer.QuestionID,
		UserID:          userID,
		RatingID:        ratingID,
		InitialAnswerID: questionAnswer.InitialAnswerID,
		FinalAnswerID:   questionAnswer.FinalAnswerID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	sql = "INSERT INTO student_answers (rationale, question_id, user_id, rating_id, initial_answer_id, final_answer_id, created_at, updated_at) VALUES (:rationale, :question_id, :user_id, :rating_id, :initial_answer_id, :final_answer_id, :created_at, :updated_at)"
	_, err = repo.DBClient.NamedExecContext(ctx, sql, studentAnswer)
	if err != nil {
		return err
	}
	return nil
}