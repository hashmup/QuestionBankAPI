package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type UserRepository interface {
	PostUserRegister(ctx context.Context, name, email, password, accountType string, schoolID int64) (*entity.UserResponse, error)
	UserLogin(ctx context.Context, email, password string) (*entity.User, error)
	GetUser(ctx context.Context, userID int64) (*entity.UserResponse, error)
}
