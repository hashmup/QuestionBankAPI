package application

import (
	"context"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
)

type questionService struct {
	questionRepo repository.QuestionRepository
}

// NewSchoolService creates a handling event service with necessary dependencies.
func NewQuestionService(questionRepo repository.QuestionRepository) service.QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (s *questionService) GetQuestionRationales(ctx context.Context, questionID, classID int64) ([]*entity.QuestionRationale, error) {
	rationales, err := s.questionRepo.GetQuestionRationales(ctx, questionID, classID)
	if err != nil {
		return nil, err
	}
	return rationales, nil
}

func (s *questionService) PostQuestions(ctx context.Context, userID int64, questionAnswer *entity.QuestionAnswer) (bool, error) {
	err := s.questionRepo.PostQuestions(ctx, userID, questionAnswer)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *questionService) GetQuestionAnswer(ctx context.Context, userID, questionID int64) (*entity.StudentAnswerResponse, error) {
	questionAnswer, err := s.questionRepo.GetQuestionAnswer(ctx, userID, questionID)
	if err != nil {
		return nil, err
	}
	return questionAnswer, nil
}

func (s *questionService) SearchQuestions(ctx context.Context, name, tag string) ([]*entity.QuestionRequest, error) {
	questions, err := s.questionRepo.SearchQuestions(ctx, name, tag)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *questionService) AnalyzeQuestion(ctx context.Context, questionID int64) (*entity.QuestionAnalysis, error) {
	questionAnalysis, err := s.questionRepo.AnalyzeQuestion(ctx, questionID)
	if err != nil {
		return nil, err
	}
	return questionAnalysis, nil
}
