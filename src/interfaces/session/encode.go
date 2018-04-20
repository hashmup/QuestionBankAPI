package session

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type PostSessionLoginResponsePayload struct {
	Session *entity.Session `json:"session"`
}

type GetSessionIsValidResponsePayload struct {
	IsValid bool `json:"is_valid"`
}

func encodePostSessionLoginResponse(obj *entity.Session) *PostSessionLoginResponsePayload {
	if obj == nil {
		obj = &entity.Session{}
	}

	payload := PostSessionLoginResponsePayload{
		Session: obj,
	}
	return &payload
}

func encodeGetSessionIsValidResponse(isValid bool) *GetSessionIsValidResponsePayload {
	payload := GetSessionIsValidResponsePayload{
		IsValid: isValid,
	}
	return &payload
}
