package school

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type GetSchoolsRequestPayload struct {
	UserID string
}

func decodeGetSchoolsRequest(r *http.Request) (*GetSchoolsRequestPayload, error) {
	payload := GetSchoolsRequestPayload{}
	payload.UserID = chi.URLParam(r, "id")
	err := interfaces.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
