package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type ClassRepository interface {
	GetClassesBySchoolID(ctx context.Context, schoolID int64) ([]*entity.Class, error)
	GetClassesByUserID(ctx context.Context, userID int64) ([]*entity.Class, error)
	GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error)
	PostFolders(ctx context.Context, classID int64, name string) error
}
