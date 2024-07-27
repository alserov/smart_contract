package middleware

import (
	"github.com/alserov/smart_contract/internal/utils"
	"net/http"
)

func WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			msg, st := utils.FromError(err)
			w.WriteHeader(st)
			w.Write([]byte(msg))
			return
		}
	}
}
