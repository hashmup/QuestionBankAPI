package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type ClassService interface {
	GetClasses(ctx context.Context, userID int64) ([]*entity.Class, error)
	JoinClass(ctx context.Context, userID, classID int64) (bool, error)
	CreateClass(ctx context.Context, userID int64, name, code, term string) (bool, error)
	GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error)
	PostFolders(ctx context.Context, classID int64, name, description string) (bool, error)
}
