package school

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetClassesRequestPayload struct {
	SchoolID int64
}

func decodeGetClassesRequest(r *http.Request) (*GetClassesRequestPayload, error) {
	payload := GetClassesRequestPayload{}
	payload.SchoolID, _ = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	err := interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
