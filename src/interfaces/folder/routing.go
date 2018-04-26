package folder

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeFolderHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getQuestionsHandler := interfaces.CustomHandler{Impl: d.GetQuestionsHandler}
	postQuestionsHandler := interfaces.CustomHandler{Impl: d.PostQuestionsHandler}

	r.Method("GET", "/folders/questions", getQuestionsHandler)
	r.Method("POST", "/folders/questions", postQuestionsHandler)
	return r
}
