package class

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeClassHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	getClassesHandler := interfaces.CustomHandler{Impl: d.GetClassesHandler}
	postInstructorClassesHandler := interfaces.CustomHandler{Impl: d.PostClassesInstructorHandler}
	postStudentClassesHandler := interfaces.CustomHandler{Impl: d.PostClassesStudentHandler}
	getFoldersHandler := interfaces.CustomHandler{Impl: d.GetFoldersHandler}
	postFoldersHandler := interfaces.CustomHandler{Impl: d.PostFoldersHandler}

	r.Method("GET", "/classes", getClassesHandler)
	r.Method("POST", "/classes/instructor", postInstructorClassesHandler)
	r.Method("POST", "/classes/student", postStudentClassesHandler)
	r.Method("GET", "/classes/folders", getFoldersHandler)
	r.Method("POST", "/classes/folders", postFoldersHandler)
	return r
}
