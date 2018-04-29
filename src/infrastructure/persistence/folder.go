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
		tags := []entity.Tag{}
		sql, args, _ = sq.Select("tags.*").From("tags").Join("question_tag_assoc ON question_tag_assoc.tag_id = tags.id").Where(sq.Eq{"question_tag_assoc.question_id": questions[i].ID}).ToSql()
		err = repo.DBClient.SelectContext(ctx, &tags, sql, args...)
		if err != nil {
			return nil, err
		}
		_answers := []*entity.Answer{}
		_tags := []string{}
		for j := range answers {
			_answers = append(_answers, &answers[j])
		}
		for j := range tags {
			_tags = append(_tags, tags[j].Name)
		}
		studentAnswer := []entity.StudentAnswer{}
		sql, args, _ = sq.Select("*").From("student_answers").Where(sq.Eq{"user_id": userID, "question_id": questions[i].ID}).ToSql()
		repo.DBClient.SelectContext(ctx, &studentAnswer, sql, args...)
		questionRequests = append(questionRequests, &entity.QuestionRequest{
			QuestionID:      questions[i].ID,
			Question:        questions[i].Question,
			Solved:          len(studentAnswer) > 0,
			Answers:         _answers,
			Tags:            _tags,
			CorrectAnswerID: questions[i].CorrectAnswerID,
		})
	}
	return questionRequests, nil
}

func (repo *folderRepository) PostQuestions(ctx context.Context, instructorID, folderID int64, questionRequest *entity.QuestionRequest) error {
	now := time.Now()
	sql := "INSERT INTO answers (name, created_at, updated_at) VALUES(:name, :created_at, :updated_at)"
	answerIDs := []int64{}
	for i := range questionRequest.Answers {
		questionRequest.Answers[i].CreatedAt = now
		questionRequest.Answers[i].UpdatedAt = now
		res, err := repo.DBClient.NamedExecContext(ctx, sql, questionRequest.Answers[i])
		if err != nil {
			return err
		}
		answerID, _ := res.LastInsertId()
		answerIDs = append(answerIDs, answerID)
	}

	tags := []entity.Tag{}
	sql, args, _ := sq.Select("*").From("tags").Where(sq.Eq{"name": questionRequest.Tags}).ToSql()
	err := repo.DBClient.SelectContext(ctx, &tags, sql, args...)
	if err != nil {
		return err
	}

	sql = "INSERT INTO tags (name, created_at, updated_at) VALUES(:name, :created_at, :updated_at)"
	tagIDs := []int64{}
	for i := range tags {
		questionRequest.Tags = remove(questionRequest.Tags, tags[i].Name)
	}
	for i := range questionRequest.Tags {
		tag := entity.Tag{
			Name:      questionRequest.Tags[i],
			CreatedAt: now,
			UpdatedAt: now,
		}
		res, err := repo.DBClient.NamedExecContext(ctx, sql, tag)
		if err != nil {
			return err
		}
		tagID, _ := res.LastInsertId()
		tagIDs = append(tagIDs, tagID)
	}
	question := entity.Question{
		Question:        questionRequest.Question,
		FolderID:        folderID,
		InstructorID:    instructorID,
		Answer1:         answerIDs[0],
		Answer2:         answerIDs[1],
		Answer3:         answerIDs[2],
		Answer4:         answerIDs[3],
		CorrectAnswerID: questionRequest.CorrectAnswerID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	sql = "INSERT INTO questions (question, folder_id, instructor_id, answer_1, answer_2, answer_3, answer_4, correct_answer_id, created_at, updated_at) VALUES (:question, :folder_id, :instructor_id, :answer_1, :answer_2, :answer_3, :answer_4, :correct_answer_id, :created_at, :updated_at)"
	res, err := repo.DBClient.NamedExecContext(ctx, sql, question)
	if err != nil {
		return err
	}
	questionID, _ := res.LastInsertId()
	sql = "INSERT INTO question_tag_assoc (question_id, tag_id, created_at, updated_at) VALUES (:question_id, :tag_id, :created_at, :updated_at)"
	for i := range tagIDs {
		questionTagAssoc := entity.QuestionTagAssoc{
			QuestionID: questionID,
			TagID:      tagIDs[i],
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		_, err := repo.DBClient.NamedExecContext(ctx, sql, questionTagAssoc)
		if err != nil {
			return err
		}
	}
	return nil
}

func remove(s []string, e string) []string {
	ret := []string{}
	for _, a := range s {
		if a != e {
			ret = append(ret, a)
		}
	}
	return ret
}
