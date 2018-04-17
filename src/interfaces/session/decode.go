package session

import (
	"encoding/json"
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type PostSessionLoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostSessionLogoutRequestPayload struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type PostSessionIsValidRequestPayload struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func decodePostSessionLoginRequest(r *http.Request) (*PostSessionLoginRequestPayload, error) {
	payload := PostSessionLoginRequestPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	err = interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func decodePostSessionLogoutRequest(r *http.Request) (*PostSessionLoginRequestPayload, error) {
	payload := PostSessionLoginRequestPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	err = interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
