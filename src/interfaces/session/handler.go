package session

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	SessionService service.SessionService
}

func (d *Dependency) PostSessionLoginHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePostSessionLoginRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	session, err := d.SessionService.PostSessionLogin(r.Context(), payload.Email, payload.Password)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodePostSessionLoginResponse(session)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}
