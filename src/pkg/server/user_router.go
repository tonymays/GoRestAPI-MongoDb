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

type userRouter struct {
	data data.Data
}

func NewUserRouter(router *mux.Router, data data.Data) *mux.Router {
	userRouter :=  userRouter{data}

	router.HandleFunc("/users", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users", VerifyToken(userRouter.createUser, data)).Methods("POST")
	router.HandleFunc("/users", VerifyToken(userRouter.getActiveUsers, data)).Methods("GET")

	router.HandleFunc("/users/{id}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.getUser, data)).Methods("GET")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.updateUser, data)).Methods("PATCH")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.activateUser, data)).Methods("PUT")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.deactivateUser, data)).Methods("DELETE")

	router.HandleFunc("/users/{id}/roles", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles", VerifyToken(userRouter.getUserRoles, data)).Methods("GET")

	router.HandleFunc("/users/{id}/roles/{roleUuid}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.assignUserRole, data)).Methods("POST")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.activateUserRole, data)).Methods("PATCH")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.deactivateUserRole, data)).Methods("PUT")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.unassignUserRole, data)).Methods("DELETE")

	return router
}

func (rcvr *userRouter) createUser(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) getActiveUsers(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) getUser(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) activateUser(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) deactivateUser(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) getUserRoles(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) assignUserRole(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) activateUserRole(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) deactivateUserRole(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *userRouter) unassignUserRole(w http.ResponseWriter, r *http.Request) {

}
