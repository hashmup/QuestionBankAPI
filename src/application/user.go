package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) PostUserRegister(ctx context.Context, name, email, password, accountType string, schoolID int64) (*entity.User, error) {
	user, err := s.userRepo.PostUserRegister(ctx, name, email, password, accountType, schoolID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
