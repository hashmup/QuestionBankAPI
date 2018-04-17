package session

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeSessionHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	postSessionLoginHandler := interfaces.CustomHandler{Impl: d.PostSessionLoginHandler}
	r.Method("POST", "/sessions/login", postSessionLoginHandler)
	return r
}
