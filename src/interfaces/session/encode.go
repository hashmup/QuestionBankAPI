package session

import "github.com/hashmup/QuestionBankAPI/src/domain/entity"

type PostSessionLoginResponsePayload struct {
	Session *entity.Session `json:"session"`
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
