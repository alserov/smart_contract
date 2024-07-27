package middleware

import (
	"fmt"
	"net/http"
)

func WithRecovery(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("panic recovery: " + p.(string))
			}
		}()

		fn(w, r)
	}
}
