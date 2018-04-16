package school

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeSchoolHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getSchoolsHandler := interfaces.CustomHandler{Impl: d.GetSchoolsHandler}
	r.Method("GET", "/schools", getSchoolsHandler)
	return r
}
