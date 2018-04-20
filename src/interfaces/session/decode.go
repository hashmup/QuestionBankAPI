package session

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type PostSessionLoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func decodeSessionHeaderRequest(r *http.Request) (*entity.SessionHeader, error) {
	payload := entity.SessionHeader{}
	auth := r.Header.Get("AuthToken")
	authInfo := strings.Split(auth, ":")
	payload.UserID, _ = strconv.ParseInt(authInfo[0], 10, 64)
	payload.Token = authInfo[1]

	err := interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
