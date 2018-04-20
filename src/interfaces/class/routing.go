package class

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeClassHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getClassesHandler := interfaces.CustomHandler{Impl: d.GetClassesHandler}
	postClassesHandler := interfaces.CustomHandler{Impl: d.PostClassesHandler}
	getFoldersHandler := interfaces.CustomHandler{Impl: d.GetFoldersHandler}
	postFoldersHandler := interfaces.CustomHandler{Impl: d.PostFoldersHandler}

	r.Method("GET", "/classes", getClassesHandler)
	r.Method("POST", "/classes", postClassesHandler)
	r.Method("GET", "/classes/folders", getFoldersHandler)
	r.Method("POST", "/classes/folders", postFoldersHandler)
	return r
}
