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

// ---- roleRouter ----
type roleRouter struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	roleService	root.RoleService
}

// ---- NewUserRouter ----
func NewRoleRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, roleService root.RoleService) *mux.Router {
	roleRouter :=  roleRouter{config, dbClient, roleService}

	router.HandleFunc("/roles", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/roles", VerifyToken(roleRouter.createRole, config, dbClient)).Methods("POST")
	router.HandleFunc("/roles", VerifyToken(roleRouter.findActiveRole, config, dbClient)).Methods("GET")

	router.HandleFunc("/roles/{id}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/roles/{id}", VerifyToken(roleRouter.findRole, config, dbClient)).Methods("GET")
	router.HandleFunc("/roles/{id}", VerifyToken(roleRouter.updateRole, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/roles/{id}", VerifyToken(roleRouter.activateRole, config, dbClient)).Methods("PUT")
	router.HandleFunc("/roles/{id}", VerifyToken(roleRouter.deactivateRole, config, dbClient)).Methods("DELETE")

	return router
}

// ---- roleRouter.createRole ----
func (rcvr *roleRouter) createRole(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		throw(w, err)
	}
	var role root.Role
	err = json.Unmarshal(body, &role)
	if err != nil {
		throw(w, err)
	}
	ro, err := rcvr.roleService.CreateRole(role)
	if err == nil {
		rcvr.respond(w, ro)
	} else {
		throw(w, err)
	}
}

// ---- roleRouter.findActiveRole ----
func (rcvr *roleRouter) findActiveRole(w http.ResponseWriter, r *http.Request) {
	var role root.Role
	role.Active = "Yes"
	roles, err := rcvr.roleService.FindRole(role)
	if err == nil {
		rcvr.respondSlice(w,roles)
	} else {
		throw(w,err)
	}
}

// ---- roleRouter.findRole ----
func (rcvr *roleRouter) findRole(w http.ResponseWriter, r *http.Request) {
	var role root.Role
	vars := mux.Vars(r)
	role.RoleId = vars["id"]
	roles, err := rcvr.roleService.FindRole(role)
	if err == nil {
		rcvr.respond(w,roles[0])
	} else {
		throw(w,err)
	}
}

// ---- roleRouter.updateRole ----
func (rcvr *roleRouter) updateRole(w http.ResponseWriter, r *http.Request) {
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
	var update root.Role
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
	var filter root.Role
	vars := mux.Vars(r)
	filter.RoleId = vars["id"]
	role, err := rcvr.roleService.UpdateRole(filter,update)
	if err == nil {
		rcvr.respond(w,role)
	} else {
		throw(w,err)
	}
}

// ---- roleRouter.activateRole ----
func (rcvr *roleRouter) activateRole(w http.ResponseWriter, r *http.Request) {
	var f root.Role
	vars := mux.Vars(r)
	f.RoleId = vars["id"]
	var u root.Role
	u.Active = "Yes"
	role, err := rcvr.roleService.UpdateRole(f,u)
	if err == nil {
		rcvr.respond(w,role)
	} else {
		throw(w,err)
	}
}

// ---- roleRouter.deactivateRole ----
func (rcvr *roleRouter) deactivateRole(w http.ResponseWriter, r *http.Request) {
	var f root.Role
	vars := mux.Vars(r)
	f.RoleId = vars["id"]
	var u root.Role
	u.Active = "No"
	role, err := rcvr.roleService.UpdateRole(f,u)
	if err == nil {
		rcvr.respond(w,role)
	} else {
		throw(w,err)
	}
}

// ---- roleRouter.respond ----
func (rcvr *roleRouter) respond(w http.ResponseWriter, r root.Role) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		throw(w,err)
	}
}
// ---- roleRouter.respondSlice ----
func (rcvr *roleRouter) respondSlice(w http.ResponseWriter, r []root.Role) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		throw(w,err)
	}
}
