package school

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type GetSchoolsResponsePayload struct {
	Schools []*entity.School `json:"schools"`
}

type GetClassesResponsePayload struct {
	Classes []*entity.Class `json:"classes"`
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

func encodeGetClassesResponse(objs []*entity.Class) *GetClassesResponsePayload {
	if objs == nil {
		objs = []*entity.Class{}
	}

	payload := GetClassesResponsePayload{
		Classes: objs,
	}
	return &payload
}
