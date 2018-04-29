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

	// res := encodeGetClassesResponse(classes)

	interfaces.Redererer.JSON(w, http.StatusOK, classes)
}

func (d *Dependency) PostClassesStudentHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadClass, err := decodePostStudentRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	succeeded, err := d.ClassService.JoinClass(r.Context(), payload.UserID, payloadClass.ClassID)
	if !succeeded || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodeGetClassesResponse(classes)

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
}

func (d *Dependency) PostClassesInstructorHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadClass, err := decodePostInstructorRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	succeeded, err := d.ClassService.CreateClass(r.Context(), payload.UserID, payloadClass.Name, payloadClass.Code, payloadClass.Term)
	if !succeeded || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodeGetClassesResponse(classes)

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
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

	// res := encodeGetFoldersResponse(folders)

	interfaces.Redererer.JSON(w, http.StatusOK, folders)
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
	succeed, err := d.ClassService.PostFolders(r.Context(), payloadFolder.ClassID, payloadFolder.Name, payloadFolder.Description)
	if !succeed || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodePostFoldersResponse(succeed)

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
}
