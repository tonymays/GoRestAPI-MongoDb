package server

import (
	"net/http"
	"pkg/data"
)

func VerifyToken(next http.HandlerFunc, data data.Data) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}