package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type folderService struct {
	folderRepo repository.FolderRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewFolderService(folderRepo repository.FolderRepository) service.FolderService {
	return &folderService{
		folderRepo: folderRepo,
	}
}

func (s *folderService) GetQuestions(ctx context.Context, userID, folderID int64) ([]*entity.QuestionRequest, error) {
	questions, err := s.folderRepo.GetQuestions(ctx, userID, folderID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *folderService) PostQuestions(ctx context.Context, profID, folderID int64, question *entity.QuestionRequest) (bool, error) {
	err := s.folderRepo.PostQuestions(ctx, profID, folderID, question)
	if err != nil {
		return false, err
	}
	return true, nil
}
