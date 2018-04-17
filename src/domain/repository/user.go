package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type UserRepository interface {
	PostUserRegister(ctx context.Context, name, email, password, accountType string, schoolID int64) (*entity.User, error)
}
