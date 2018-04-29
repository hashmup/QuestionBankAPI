package folder

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	FolderService  service.FolderService
	SessionService service.SessionService
}

func (d *Dependency) GetQuestionsHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadQuestion, err := decodeGetQuestionsRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	questions, err := d.FolderService.GetQuestions(r.Context(), payload.UserID, payloadQuestion.FolderID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodeGetQuestionsResponse(questions)

	interfaces.Redererer.JSON(w, http.StatusOK, questions)
}

func (d *Dependency) PostQuestionsHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadFolder, err := decodePostQuestionsRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	succeed, err := d.FolderService.PostQuestions(r.Context(), payload.UserID, payloadFolder.FolderID, &entity.QuestionRequest{
		Question:        payloadFolder.Question,
		Answers:         payloadFolder.Answers,
		Tags:            payloadFolder.Tags,
		CorrectAnswerID: payloadFolder.CorrectAnswerID,
	})
	if !succeed || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
}
