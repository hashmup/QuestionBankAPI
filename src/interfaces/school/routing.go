package school

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeSchoolHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getSchoolsHandler := interfaces.CustomHandler{Impl: d.GetSchoolsHandler}
	getClassesHandler := interfaces.CustomHandler{Impl: d.GetClassesHandler}
	r.Method("GET", "/schools", getSchoolsHandler)
	r.Method("GET", "/schools/{id}/classes", getClassesHandler)
	return r
}
