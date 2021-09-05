package server

import (
	"net/http"
	"pkg/data"
)

func VerifyToken(next http.HandlerFunc, config data.Configuration) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}