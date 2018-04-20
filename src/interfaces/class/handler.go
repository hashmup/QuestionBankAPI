package class

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	ClassService   service.ClassService
	SessionService service.SessionService
}

func (d *Dependency) GetClassesHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeSessionHeaderRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	isValid, err := d.SessionService.IsValidSession(r.Context(), payload.UserID, payload.Token)
	if !isValid || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	classes, err := d.ClassService.GetClasses(r.Context(), payload.UserID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodeGetClassesResponse(classes)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}

func (d *Dependency) PostClassesHandler(w http.ResponseWriter, r *http.Request) {
}

func (d *Dependency) GetFoldersHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeSessionHeaderRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	isValid, err := d.SessionService.IsValidSession(r.Context(), payload.UserID, payload.Token)
	if !isValid || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	payloadFolder, err := decodeGetFoldersRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	folders, err := d.ClassService.GetFolders(r.Context(), payloadFolder.ClassID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodeGetFoldersResponse(folders)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}

func (d *Dependency) PostFoldersHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := decodeSessionHeaderRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	isValid, err := d.SessionService.IsValidSession(r.Context(), payload.UserID, payload.Token)
	if !isValid || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	payloadFolder, err := decodePostFoldersRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	succeed, err := d.ClassService.PostFolders(r.Context(), payloadFolder.ClassID, payloadFolder.Name)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodePostFoldersResponse(succeed)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}
