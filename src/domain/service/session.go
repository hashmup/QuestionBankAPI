package service

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
)

type SessionService interface {
	PostSessionLogin(ctx context.Context, email, password string) (*entity.Session, error)
}
