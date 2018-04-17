package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type ClassRepository interface {
	GetClasses(ctx context.Context, schoolID int64) ([]*entity.Class, error)
}
