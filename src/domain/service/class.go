package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type ClassService interface {
	GetClasses(ctx context.Context, userID int64) ([]*entity.Class, error)
	GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error)
	PostFolders(ctx context.Context, classID int64, name string) (bool, error)
}
