package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type FolderService interface {
	GetQuestions(ctx context.Context, userID, folderID int64) ([]*entity.QuestionRequest, error)
	PostQuestions(ctx context.Context, profID, folderID int64, question *entity.QuestionRequest) (bool, error)
}
