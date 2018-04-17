package user

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type PostUserRegisterResponsePayload struct {
	User *entity.User `json:"user"`
}

func encodePostUserRegisterResponse(obj *entity.User) *PostUserRegisterResponsePayload {
	if obj == nil {
		obj = &entity.User{}
	}

	payload := PostUserRegisterResponsePayload{
		User: obj,
	}
	return &payload
}
