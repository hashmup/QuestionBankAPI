package user

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeUserHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	postUserRegisterHandler := interfaces.CustomHandler{Impl: d.PostUserRegisterHandler}
	getUsersHandler := interfaces.CustomHandler{Impl: d.GetUsersHandler}
	r.Method("POST", "/users/register", postUserRegisterHandler)
	r.Method("GET", "/users", getUsersHandler)
	return r
}
