package school

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetClassesRequestPayload struct {
	SchoolID  int64
	Name      string
	ClassCode string
	Term      string
}

func decodeGetClassesRequest(r *http.Request) (*GetClassesRequestPayload, error) {
	payload := GetClassesRequestPayload{}
	name := r.URL.Query().Get("name")
	classCode := r.URL.Query().Get("class_code")
	term := r.URL.Query().Get("term")
	payload.SchoolID, _ = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	payload.Name = name
	payload.ClassCode = classCode
	payload.Term = term
	err := interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
