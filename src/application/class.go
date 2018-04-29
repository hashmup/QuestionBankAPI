package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type classService struct {
	classRepo repository.ClassRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewClassService(classRepo repository.ClassRepository) service.ClassService {
	return &classService{
		classRepo: classRepo,
	}
}

func (s *classService) GetClasses(ctx context.Context, userID int64) ([]*entity.Class, error) {
	classes, err := s.classRepo.GetClassesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *classService) GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error) {
	folders, err := s.classRepo.GetFolders(ctx, classID)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func (s *classService) PostFolders(ctx context.Context, classID int64, name, description string) (bool, error) {
	err := s.classRepo.PostFolders(ctx, classID, name, description)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *classService) JoinClass(ctx context.Context, userID, classID int64) (bool, error) {
	err := s.classRepo.JoinClass(ctx, userID, classID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *classService) CreateClass(ctx context.Context, userID int64, name, code, term string) (bool, error) {
	err := s.classRepo.CreateClass(ctx, userID, name, code, term)
	if err != nil {
		return false, err
	}
	return true, nil
}
