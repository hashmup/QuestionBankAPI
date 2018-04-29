package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type SchoolService interface {
	GetSchools(ctx context.Context) ([]*entity.School, error)
	GetClasses(ctx context.Context, schoolID int64, name, classCode, term string) ([]*entity.Class, error)
}
