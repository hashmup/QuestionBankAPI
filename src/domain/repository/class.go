package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type ClassRepository interface {
	GetClassesBySchoolID(ctx context.Context, schoolID int64) ([]*entity.Class, error)
	GetClasses(ctx context.Context, schoolID int64, name, classCode, term string) ([]*entity.Class, error)
	GetClassesByUserID(ctx context.Context, userID int64) ([]*entity.Class, error)
	GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error)
	PostFolders(ctx context.Context, classID int64, name, description string) error
	JoinClass(ctx context.Context, userID, classID int64) error
	CreateClass(ctx context.Context, userID int64, name, code, term string) error
}
