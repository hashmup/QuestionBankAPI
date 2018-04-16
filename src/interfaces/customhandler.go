package interfaces

import (
	"net/http"
)

type CustomHandler struct {
	Impl func(http.ResponseWriter, *http.Request)
}

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Impl(w, r)
}
