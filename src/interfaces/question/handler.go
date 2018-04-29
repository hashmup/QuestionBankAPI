package question

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	QuestionService service.QuestionService
	SessionService  service.SessionService
}

func (d *Dependency) GetQuestionRationaleHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadQuestion, err := decodeGetQuestionRationaleRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	rationales, err := d.QuestionService.GetQuestionRationales(r.Context(), payloadQuestion.QuestionID, payloadQuestion.ClassID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodeGetQuestionsResponse(questions)

	interfaces.Redererer.JSON(w, http.StatusOK, rationales)
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
	payloadQuestion, err := decodePostQuestionsRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	succeed, err := d.QuestionService.PostQuestions(r.Context(), payload.UserID, &entity.QuestionAnswer{
		QuestionID:      payloadQuestion.QuestionID,
		Rationale:       payloadQuestion.Rationale,
		InitialAnswerID: payloadQuestion.InitialAnswerID,
		FinalAnswerID:   payloadQuestion.FinalAnswerID,
		Rating:          payloadQuestion.Rating,
	})
	if !succeed || err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	interfaces.Redererer.JSON(w, http.StatusOK, nil)
}

func (d *Dependency) GetQuestionAnswerHandler(w http.ResponseWriter, r *http.Request) {
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
	payloadQuestion, err := decodeGetQuestionAnswerRequest(r)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusBadRequest, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}
	answers, err := d.QuestionService.GetQuestionAnswer(r.Context(), payload.UserID, payloadQuestion.QuestionID)
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	// res := encodeGetQuestionsResponse(questions)

	interfaces.Redererer.JSON(w, http.StatusOK, answers)
}
