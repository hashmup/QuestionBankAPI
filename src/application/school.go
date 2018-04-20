package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type schoolService struct {
	schoolRepo repository.SchoolRepository
	classRepo  repository.ClassRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewSchoolService(schoolRepo repository.SchoolRepository, classRepo repository.ClassRepository) service.SchoolService {
	return &schoolService{
		schoolRepo: schoolRepo,
		classRepo:  classRepo,
	}
}

func (s *schoolService) GetSchools(ctx context.Context) ([]*entity.School, error) {
	schools, err := s.schoolRepo.GetSchools(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (s *schoolService) GetClasses(ctx context.Context, schoolID int64) ([]*entity.Class, error) {
	classes, err := s.classRepo.GetClassesBySchoolID(ctx, schoolID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
