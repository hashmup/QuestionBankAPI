package user

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeUserHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	postUserRegisterHandler := interfaces.CustomHandler{Impl: d.PostUserRegisterHandler}
	r.Method("POST", "/users/register", postUserRegisterHandler)
	return r
}
