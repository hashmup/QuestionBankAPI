package question

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetQuestionRationaleRequestPayload struct {
	QuestionID int64 `json:"question_id"`
	ClassID    int64 `json:"class_id"`
}

type GetQuestionAnswerRequestPayload struct {
	QuestionID int64 `json:"question_id"`
}

type PostQuestionsRequestPayload struct {
	QuestionID      int64          `json:"question_id"`
	Rationale       string         `json:"rationale"`
	InitialAnswerID int64          `json:"initial_answer_id"`
	FinalAnswerID   int64          `json:"final_answer_id"`
	Rating          *entity.Rating `json:"rating"`
}

type SearchQuestionRequestPayload struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
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

func decodeGetQuestionRationaleRequest(r *http.Request) (*GetQuestionRationaleRequestPayload, error) {
	payload := GetQuestionRationaleRequestPayload{}
	questionID, err := strconv.ParseInt(r.URL.Query().Get("question_id"), 10, 64)
	classID, err := strconv.ParseInt(r.URL.Query().Get("class_id"), 10, 64)
	if err != nil {
		return nil, err
	}
	payload.QuestionID = questionID
	payload.ClassID = classID

	return &payload, nil
}

func decodeGetQuestionAnswerRequest(r *http.Request) (*GetQuestionAnswerRequestPayload, error) {
	payload := GetQuestionAnswerRequestPayload{}
	questionID, err := strconv.ParseInt(r.URL.Query().Get("question_id"), 10, 64)
	if err != nil {
		return nil, err
	}
	payload.QuestionID = questionID

	return &payload, nil
}

func decodeSearchQuestionRequest(r *http.Request) (*SearchQuestionRequestPayload, error) {
	payload := SearchQuestionRequestPayload{}
	name := r.URL.Query().Get("name")
	tag := r.URL.Query().Get("tag")
	payload.Name = name
	payload.Tag = tag

	return &payload, nil
}
