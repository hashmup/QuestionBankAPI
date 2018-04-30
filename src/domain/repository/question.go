package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type QuestionRepository interface {
	GetQuestionRationales(ctx context.Context, questionID, classID int64) ([]*entity.QuestionRationale, error)
	GetQuestionAnswer(ctx context.Context, userID, questionID int64) (*entity.StudentAnswerResponse, error)
	PostQuestions(ctx context.Context, userID int64, questionAnswer *entity.QuestionAnswer) error
	SearchQuestions(ctx context.Context, name, tag string) ([]*entity.QuestionRequest, error)
	AnalyzeQuestion(ctx context.Context, questionID int64) (*entity.QuestionAnalysis, error)
}
