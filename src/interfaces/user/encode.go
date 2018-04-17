package user

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type PostUserRegisterResponsePayload struct {
	User *entity.User `json:"user"`
}

type PostUserLoginResponsePayload struct {
	Session *entity.Session `json:"session"`
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

func encodePostUserLoginResponse(obj *entity.Session) *PostUserLoginResponsePayload {
	if obj == nil {
		obj = &entity.Session{}
	}

	payload := PostUserLoginResponsePayload{
		Session: obj,
	}
	return &payload
}
