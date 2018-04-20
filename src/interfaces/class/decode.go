package class

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetFoldersRequestPayload struct {
	ClassID int64
}

type PostFoldersRequestPayload struct {
	ClassID int64  `json:"class_id"`
	Name    string `json:"name"`
}

func decodeGetFoldersRequest(r *http.Request) (*GetFoldersRequestPayload, error) {
	payload := GetFoldersRequestPayload{}
	classID, err := strconv.ParseInt(r.URL.Query().Get("class_id"), 10, 64)
	if err != nil {
		return nil, err
	}
	payload.ClassID = classID

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

func decodePostFoldersRequest(r *http.Request) (*PostFoldersRequestPayload, error) {
	payload := PostFoldersRequestPayload{}
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