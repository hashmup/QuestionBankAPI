package folder

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetQuestionsRequestPayload struct {
	FolderID int64
}

type PostQuestionsRequestPayload struct {
	FolderID        int64            `json:"folder_id"`
	Question        string           `json:"question"`
	Answers         []*entity.Answer `json:"answers"`
	CorrectAnswerID int64            `json:"correct_answer_id"`
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

func decodePostQuestionsRequest(r *http.Request) (*PostQuestionsRequestPayload, error) {
	payload := PostQuestionsRequestPayload{}
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

func decodeGetQuestionsRequest(r *http.Request) (*GetQuestionsRequestPayload, error) {
	payload := GetQuestionsRequestPayload{}
	folderID, err := strconv.ParseInt(r.URL.Query().Get("folder_id"), 10, 64)
	if err != nil {
		return nil, err
	}
	payload.FolderID = folderID

	return &payload, nil
}
