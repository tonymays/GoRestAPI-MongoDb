package server

import (
	"net/http"
	"pkg/configuration"
)

// ---- VerifyToken ----
func VerifyToken(next http.HandlerFunc, config configuration.Configuration) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}