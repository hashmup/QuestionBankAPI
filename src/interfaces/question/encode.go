package question

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type GetQuestionsResponsePayload struct {
	Questions []*entity.QuestionRequest `json:"questions"`
}

type GetFoldersResponsePayload struct {
	Folders []*entity.Folder `json:"folders"`
}

func encodeGetQuestionsResponse(objs []*entity.QuestionRequest) *GetQuestionsResponsePayload {
	if objs == nil {
		objs = []*entity.QuestionRequest{}
	}

	payload := GetQuestionsResponsePayload{
		Questions: objs,
	}
	return &payload
}
