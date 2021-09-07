package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"net/http"
	"pkg"
	"pkg/configuration"
)

// ---- userRouter ----
type userRouter struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	userService	root.UserService
}

// ---- NewUserRouter ----
func NewUserRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, userService root.UserService) *mux.Router {
	userRouter :=  userRouter{config, dbClient, userService}

	router.HandleFunc("/users", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users", VerifyToken(userRouter.createUser, config, dbClient)).Methods("POST")
	router.HandleFunc("/users", VerifyToken(userRouter.findActiveUsers, config, dbClient)).Methods("GET")

	router.HandleFunc("/users/{id}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.findUser, config, dbClient)).Methods("GET")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.updateUser, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.activateUser, config, dbClient)).Methods("PUT")
	router.HandleFunc("/users/{id}", VerifyToken(userRouter.deactivateUser, config, dbClient)).Methods("DELETE")

	router.HandleFunc("/users/{id}/roles", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles", VerifyToken(userRouter.findUserRoles, config, dbClient)).Methods("GET")

	router.HandleFunc("/users/{id}/roles/{roleUuid}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.assignUserRole, config, dbClient)).Methods("POST")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.activateUserRole, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/users/{id}/roles/{roleUuid}", VerifyToken(userRouter.deactivateUserRole, config, dbClient)).Methods("PUT")

	return router
}

// ---- userRouter.createUser ----
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

// ---- userRouter.findActiveUsers ----
func (rcvr *userRouter) findActiveUsers(w http.ResponseWriter, r *http.Request) {
	var u root.User
	u.Active = "Yes"
	users, err := rcvr.userService.FindUser(u)
	if err == nil {
		rcvr.respondSlice(w,users)
	} else {
		throw(w,err)
	}
}

// ---- userRouter.findUser ----
func (rcvr *userRouter) findUser(w http.ResponseWriter, r *http.Request) {
	var u root.User
	vars := mux.Vars(r)
	u.UserId = vars["id"]
	users, err := rcvr.userService.FindUser(u)
	if err == nil {
		rcvr.respond(w,users[0])
	} else {
		throw(w,err)
	}
}

// ---- userRouter.updateUser ----
func (rcvr *userRouter) updateUser(w http.ResponseWriter, r *http.Request) {
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
	var update root.User
	err = json.Unmarshal(body, &update)
	if err != nil {
		throw(w,err)
		return
	}
	err = update.Validate(false)
	if err != nil {
		throw(w,err)
		return
	}
	var filter root.User
	vars := mux.Vars(r)
	filter.UserId = vars["id"]
	user, err := rcvr.userService.UpdateUser(filter,update)
	if err == nil {
		rcvr.respond(w,user)
	} else {
		throw(w,err)
	}
}

// ---- userRouter.activateUser ----
func (rcvr *userRouter) activateUser(w http.ResponseWriter, r *http.Request) {
	var f root.User
	vars := mux.Vars(r)
	f.UserId = vars["id"]
	var u root.User
	u.Active = "Yes"
	user, err := rcvr.userService.UpdateUser(f,u)
	if err == nil {
		rcvr.respond(w,user)
	} else {
		throw(w,err)
	}
}

// ---- userRouter.deactivateUser ----
func (rcvr *userRouter) deactivateUser(w http.ResponseWriter, r *http.Request) {
	var f root.User
	vars := mux.Vars(r)
	f.UserId = vars["id"]
	var u root.User
	u.Active = "No"
	user, err := rcvr.userService.UpdateUser(f,u)
	if err == nil {
		rcvr.respond(w,user)
	} else {
		throw(w,err)
	}
}

// ---- userRouter.findUserRoles ----
func (rcvr *userRouter) findUserRoles(w http.ResponseWriter, r *http.Request) {

}

// ---- userRouter.assignUserRole ----
func (rcvr *userRouter) assignUserRole(w http.ResponseWriter, r *http.Request) {

}

// ---- userRouter.activateUserRole ----
func (rcvr *userRouter) activateUserRole(w http.ResponseWriter, r *http.Request) {

}

// ---- userRouter.deactivateUserRole ----
func (rcvr *userRouter) deactivateUserRole(w http.ResponseWriter, r *http.Request) {

}

// ---- userRouter.respond ----
func (rcvr *userRouter) respond(w http.ResponseWriter, u root.User) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		throw(w,err)
	}
}
// ---- userRouter.respondSlice ----
func (rcvr *userRouter) respondSlice(w http.ResponseWriter, u []root.User) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		throw(w,err)
	}
}
