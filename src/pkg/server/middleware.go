package server

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"pkg/configuration"
)

func VerifyToken(next http.HandlerFunc, config configuration.Configuration, dbClient *mongo.Client) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}