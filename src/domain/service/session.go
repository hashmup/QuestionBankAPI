package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type SessionService interface {
	LoginSession(ctx context.Context, email, password string) (*entity.Session, error)
	LogoutSession(ctx context.Context, userID int64, token string) (bool, error)
	IsValidSession(ctx context.Context, userID int64, token string) (bool, error)
}
