package question

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeQuestionHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getQuestionRationaleHandler := interfaces.CustomHandler{Impl: d.GetQuestionRationaleHandler}
	postQuestionsHandler := interfaces.CustomHandler{Impl: d.PostQuestionsHandler}
	getQuestionAnswerHandler := interfaces.CustomHandler{Impl: d.GetQuestionAnswerHandler}
	searchQuestionHandler := interfaces.CustomHandler{Impl: d.SearchQuestionHandler}

	r.Method("GET", "/questions/rationales", getQuestionRationaleHandler)
	r.Method("POST", "/questions", postQuestionsHandler)
	r.Method("GET", "/questions/answer", getQuestionAnswerHandler)
	r.Method("GET", "/questions/search", searchQuestionHandler)
	return r
}
