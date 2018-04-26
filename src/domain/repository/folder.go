package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type FolderRepository interface {
	GetQuestions(ctx context.Context, userID, folderID int64) ([]*entity.QuestionRequest, error)
	PostQuestions(ctx context.Context, profID, classID int64, question *entity.QuestionRequest) error
}
