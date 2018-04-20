package application

import (
	"context"
	"fmt"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"golang.org/x/crypto/bcrypt"
)

type sessionService struct {
	sessionRepo repository.SessionRepository
	userRepo    repository.UserRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewSessionService(sessionRepo repository.SessionRepository, userRepo repository.UserRepository) service.SessionService {
	return &sessionService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (s *sessionService) LoginSession(ctx context.Context, email, password string) (*entity.Session, error) {
	user, err := s.userRepo.UserLogin(ctx, email, password)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		fmt.Printf("Password does not match")
		return nil, err
	}

	session, err := s.sessionRepo.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) LogoutSession(ctx context.Context, userID int64, token string) (bool, error) {
	succeed, err := s.sessionRepo.DeleteSession(ctx, userID, token)
	if err != nil {
		panic(err)
	}
	return succeed, nil
}

func (s *sessionService) IsValidSession(ctx context.Context, userID int64, token string) (bool, error) {
	isValid, err := s.sessionRepo.IsValidSession(ctx, userID, token)
	if err != nil {
		panic(err)
	}
	return isValid, nil
}
