package session

import (
	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

func MakeSessionHandler(d *Dependency, r *chi.Mux) *chi.Mux {
	postSessionLoginHandler := interfaces.CustomHandler{Impl: d.PostSessionLoginHandler}
	postSessionLogoutHandler := interfaces.CustomHandler{Impl: d.PostSessionLogoutHandler}
	getSessionIsValidHandler := interfaces.CustomHandler{Impl: d.GetSessionIsValidHandler}
	r.Method("POST", "/sessions/login", postSessionLoginHandler)
	r.Method("POST", "/sessions/logout", postSessionLogoutHandler)
	r.Method("GET", "/sessions/isvalid", getSessionIsValidHandler)
	return r
}
