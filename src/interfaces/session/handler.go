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
	session, err := d.SessionService.LoginSession(r.Context(), payload.Email, payload.Password)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodePostSessionLoginResponse(session)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}

func (d *Dependency) PostSessionLogoutHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeSessionHeaderRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	isValid, err := d.SessionService.LogoutSession(r.Context(), payload.UserID, payload.Token)
	if !isValid || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
}

func (d *Dependency) GetSessionIsValidHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeSessionHeaderRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	isValid, err := d.SessionService.IsValidSession(r.Context(), payload.UserID, payload.Token)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodeGetSessionIsValidResponse(isValid)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}
