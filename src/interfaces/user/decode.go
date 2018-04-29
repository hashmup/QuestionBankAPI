package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type PostUserRegisterRequestPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	SchoolID int64  `json:"school_id"`
	Type     string `json:"type"`
}

type PostUserLoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUsersRequestPayload struct {
	UserID int64 `json:"user_id"`
}

func decodePostUserRegisterRequest(r *http.Request) (*PostUserRegisterRequestPayload, error) {
	payload := PostUserRegisterRequestPayload{}
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

func decodePostUserLoginRequest(r *http.Request) (*PostUserLoginRequestPayload, error) {
	payload := PostUserLoginRequestPayload{}
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

func decodeGetUsersRequest(r *http.Request) (*GetUsersRequestPayload, error) {
	payload := GetUsersRequestPayload{}
	userID, _ := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	payload.UserID = userID
	err := interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
