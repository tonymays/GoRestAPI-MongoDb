package server

import (
	"encoding/json"
//	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"net/http"
	"pkg"
	"pkg/configuration"
	"time"
)

// ---- authRouter ----
type authRouter struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	authService	root.AuthService
}

// ---- NewAuthRouter ----
func NewAuthRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, authService root.AuthService) *mux.Router {
	authRouter :=  authRouter{config, dbClient, authService}
	router.HandleFunc("/auth", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/auth", authRouter.startSession).Methods("POST")
	router.HandleFunc("/auth", VerifyToken(authRouter.killSession, config, dbClient)).Methods("DELETE")
	router.HandleFunc("/auth", VerifyToken(authRouter.refreshSession, config, dbClient)).Methods("GET")
	router.HandleFunc("/auth", VerifyToken(authRouter.checkSession, config, dbClient)).Methods("HEAD")
	router.HandleFunc("/auth", VerifyToken(authRouter.changePassword, config, dbClient)).Methods("PUT")
	return router
}

// ---- authRouter.startSession ----
func (rcvr *authRouter) startSession(w http.ResponseWriter, r *http.Request) {
	var payload root.AuthPayload
	var userToken root.UserToken
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		throw(w,err)
		return
	}
	err = r.Body.Close()
	if err != nil {
		throw(w,err)
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		throw(w,err)
		return
	}
	payload.LoginIp = GetIpAddress(r)
	userToken, err = rcvr.authService.StartSession(payload)
	if err == nil {
		userToken.RemoteAddr = payload.LoginIp
		expireDate := time.Now().Add(time.Hour * 1).Unix()
		payload.AuthToken = CreateToken(userToken, rcvr.config, expireDate)
		w = SetResponseHeaders(w, payload.AuthToken, "")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userToken)
	} else {
		throw(w, err)
	}
}

// ---- authRouter.killSession ----
func (rcvr *authRouter) killSession(w http.ResponseWriter, r *http.Request) {
	// destroy the current AuthToken which will cause the middleware to reject
	// all calls until a valid AuthToken is presented
	var payload root.AuthPayload
	payload.AuthToken = r.Header.Get("Auth-Token")
	rcvr.authService.KillSession(payload)
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
}

// ---- authRouter.refreshSession ----
func (rcvr *authRouter) refreshSession(w http.ResponseWriter, r *http.Request) {
	// kill the old session
	var payload root.AuthPayload
	payload.AuthToken = r.Header.Get("Auth-Token")
	rcvr.authService.KillSession(payload)

	// create a new session
	var userToken root.UserToken
	userToken = DecodeJWT(payload.AuthToken, rcvr.config)
	expireDate := time.Now().Add(time.Hour * 1).Unix()
	userToken.RemoteAddr = GetIpAddress(r)
	payload.AuthToken = CreateToken(userToken, rcvr.config, expireDate)

	// return the results
	w = SetResponseHeaders(w, payload.AuthToken, "")
	w.WriteHeader(http.StatusOK)
}

// ---- authRouter.checkSession ----
func (rcvr *authRouter) checkSession(w http.ResponseWriter, r *http.Request) {
	// set the response headers
	w = SetResponseHeaders(w, "", "")

	// grab and validate the Auth Token
	var payload root.AuthPayload
	payload.AuthToken = r.Header.Get("Auth-Token")
	err := rcvr.authService.CheckSession(payload)

	// return ok on no errors
	if err == nil {
		w.WriteHeader(http.StatusOK)
	// else return forbidden
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// ---- authRouter.changePassword ----
func (rcvr *authRouter) changePassword(w http.ResponseWriter, r *http.Request) {
	var payload root.ChangePasswordPayload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		throw(w, err)
		return
	}
	err = r.Body.Close()
	if err != nil {
		throw(w, err)
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		throw(w, err)
		return
	}
	err = rcvr.authService.ChangePassword(payload)
	if err == nil {
		w = SetResponseHeaders(w, "", "")
		w.WriteHeader(http.StatusAccepted)
		payload.Password = ""
		payload.NewPassword = ""
		json.NewEncoder(w).Encode(payload)
	} else {
		throw(w, err)
	}
}