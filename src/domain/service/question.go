package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type QuestionService interface {
	GetQuestionRationales(ctx context.Context, questionID, classID int64) ([]*entity.QuestionRationale, error)
	GetQuestionAnswer(ctx context.Context, userID, questionID int64) (*entity.StudentAnswerResponse, error)
	PostQuestions(ctx context.Context, userID int64, questionAnswer *entity.QuestionAnswer) (bool, error)
	SearchQuestions(ctx context.Context, name, tag string) ([]*entity.QuestionRequest, error)
}
