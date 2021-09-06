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
)

type userRouter struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	userService	root.UserService
}

func NewUserRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, userService root.UserService) *mux.Router {
	userRouter :=  userRouter{config, dbClient, userService}

	router.HandleFunc("/users", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users", VerifyToken(userRouter.createUser, config, dbClient)).Methods("POST")
	router.HandleFunc("/users", VerifyToken(userRouter.getActiveUsers, config, dbClient)).Methods("GET")

	router.HandleFunc("/users/{id}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.getUser, config, dbClient)).Methods("GET")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.updateUser, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.activateUser, config, dbClient)).Methods("PUT")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.deactivateUser, config, dbClient)).Methods("DELETE")

	router.HandleFunc("/users/{id}/roles", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles", VerifyToken(userRouter.getUserRoles, config, dbClient)).Methods("GET")

	router.HandleFunc("/users/{id}/roles/{roleUuid}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.assignUserRole, config, dbClient)).Methods("POST")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.activateUserRole, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.deactivateUserRole, config, dbClient)).Methods("PUT")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.unassignUserRole, config, dbClient)).Methods("DELETE")

	return router
}

func (rcvr *userRouter) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		throw(w, err)
	}
	var user root.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		throw(w, err)
	}
	u, err := rcvr.userService.CreateUser(user)
	if err == nil {
		rcvr.respond(w, u)
	} else {
		throw(w, err)
	}


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

func (rcvr *userRouter) respond(w http.ResponseWriter, u root.User) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
