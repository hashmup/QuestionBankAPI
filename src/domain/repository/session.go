package repository

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, userID int64) (*entity.Session, error)
	DeleteSession(ctx context.Context, userID int64, token string) (bool, error)
	IsValidSession(ctx context.Context, userID int64, token string) (bool, error)
}
