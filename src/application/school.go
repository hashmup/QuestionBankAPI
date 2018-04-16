package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type schoolService struct {
	repo repository.SchoolRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewSchoolService(repo repository.SchoolRepository) service.SchoolService {
	return &schoolService{
		repo: repo,
	}
}

func (s *schoolService) GetSchools(ctx context.Context) ([]*entity.School, error) {
	schools, err := s.repo.GetSchools(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}
