package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"pkg"
	"pkg/configuration"
	"time"
)

// ---- VerifyToken ----
func VerifyToken(next http.HandlerFunc, config configuration.Configuration, dbClient *mongo.Client) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var jwtError JWTError
		var MySigningKey = []byte(config.Secret)
		authToken := r.Header.Get("Auth-Token")
		if isTokenBlacklisted(authToken, config, dbClient) {
			jwtError.Message = "invalid token"
			respondWithError(w, http.StatusUnauthorized, jwtError)
			return
		}
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error")
			}
			return []byte(MySigningKey), nil
		})
		if err != nil {
			jwtError.Message = err.Error()
			respondWithError(w, http.StatusUnauthorized, jwtError)
			return
		}
		tokenClaims := token.Claims.(jwt.MapClaims)
		userId := tokenClaims["user_id"].(string)
		remoteAddr := tokenClaims["remote_addr"].(string)
		requestIp := GetIpAddress(r)
		if remoteAddr != requestIp {
			jwtError.Message = "invalid token"
			respondWithError(w, http.StatusUnauthorized, jwtError)
			return
		}
		var u root.User
		collection := dbClient.Database(config.DbName).Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		err = collection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&u)
		if err != nil {
			jwtError.Message = "invalid token"
			respondWithError(w, http.StatusUnauthorized, jwtError)
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			jwtError.Message = "invalid token"
			respondWithError(w, http.StatusUnauthorized, jwtError)
			return
		}
	})
}

// ---- isTokenBlacklisted ----
func isTokenBlacklisted(authToken string, config configuration.Configuration, client *mongo.Client) bool {
	var b root.Blacklist
	collection := client.Database(config.DbName).Collection("blacklists")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"auth_token": authToken}).Decode(&b)
	return (err == nil)
}

// ---- respondWithError ----
func respondWithError(w http.ResponseWriter, status int, error JWTError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(error)
	if err != nil {
		panic(err)
	}
}
