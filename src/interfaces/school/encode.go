package school

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type GetSchoolsResponsePayload struct {
	Schools []*entity.School `json:"schools"`
}

func encodeGetSchoolsResponse(objs []*entity.School) *GetSchoolsResponsePayload {
	if objs == nil {
		objs = []*entity.School{}
	}

	payload := GetSchoolsResponsePayload{
		Schools: objs,
	}
	return &payload
}
