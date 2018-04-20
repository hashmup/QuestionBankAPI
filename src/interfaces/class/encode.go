package class

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type GetClassesResponsePayload struct {
	Classes []*entity.Class `json:"classes"`
}

type GetFoldersResponsePayload struct {
	Folders []*entity.Folder `json:"folders"`
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

func encodeGetFoldersResponse(objs []*entity.Folder) *GetFoldersResponsePayload {
	if objs == nil {
		objs = []*entity.Folder{}
	}

	payload := GetFoldersResponsePayload{
		Folders: objs,
	}
	return &payload
}
