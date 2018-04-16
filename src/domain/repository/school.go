package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type SchoolRepository interface {
	GetSchools(ctx context.Context) ([]*entity.School, error)
}
