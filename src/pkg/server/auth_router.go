package server

import (
//	"encoding/json"
//	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
//	"io"
//	"io/ioutil"
	"net/http"
	"pkg"
	"pkg/configuration"
)

type authRouter struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	authService	root.AuthService
}

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

func (rcvr *authRouter) startSession(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *authRouter) killSession(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *authRouter) refreshSession(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *authRouter) checkSession(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *authRouter) changePassword(w http.ResponseWriter, r *http.Request) {

}