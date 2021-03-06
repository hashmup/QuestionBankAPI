package user

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	UserService service.UserService
}

func (d *Dependency) PostUserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePostUserRegisterRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	user, err := d.UserService.PostUserRegister(r.Context(), payload.Name, payload.Email, payload.Password, payload.Type, payload.SchoolID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	interfaces.Redererer.JSON(w, http.StatusOK, user)
}

func (d *Dependency) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeGetUsersRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	user, err := d.UserService.GetUser(r.Context(), payload.UserID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	interfaces.Redererer.JSON(w, http.StatusOK, user)
}
