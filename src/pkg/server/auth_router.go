package server

import (
//	"encoding/json"
//	"errors"
	"github.com/gorilla/mux"
//	"go.mongodb.org/mongo-driver/mongo"
//	"io"
//	"io/ioutil"
	"net/http"
	"pkg/data"
)

type authRouter struct {
	data data.Data
}

func NewAuthRouter(router *mux.Router, data data.Data) *mux.Router {
	authRouter :=  authRouter{data}
	router.HandleFunc("/auth", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/auth", authRouter.startSession).Methods("POST")
	router.HandleFunc("/auth", VerifyToken(authRouter.killSession, data)).Methods("DELETE")
	router.HandleFunc("/auth", VerifyToken(authRouter.refreshSession, data)).Methods("GET")
	router.HandleFunc("/auth", VerifyToken(authRouter.checkSession, data)).Methods("HEAD")
	router.HandleFunc("/auth", VerifyToken(authRouter.changePassword, data)).Methods("PUT")
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