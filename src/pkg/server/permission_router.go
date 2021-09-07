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
type permissionRouter struct {
	config				configuration.Configuration
	dbClient			*mongo.Client
	permissionService	root.PermissionService
}

// ---- NewUserRouter ----
func NewPermissionRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, permissionService root.PermissionService) *mux.Router {
	permissionRouter :=  userRouter{config, dbClient, permissionService}

	router.HandleFunc("/permissions", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/permissions", VerifyToken(permissionRouter.createPermission, config, dbClient)).Methods("POST")
	router.HandleFunc("/permissions", VerifyToken(permissionRouter.findActivePermissions, config, dbClient)).Methods("GET")

	router.HandleFunc("/permissions/{id}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/permissions/{id}", VerifyToken(permissionRouter.findPermission, config, dbClient)).Methods("GET")
	router.HandleFunc("/permissions/{id}", VerifyToken(permissionRouter.updatePermission, config, dbClient)).Methods("PATCH")
	router.HandleFunc("/permissions/{id}", VerifyToken(permissionRouter.activatePermission, config, dbClient)).Methods("PUT")
	router.HandleFunc("/permissions/{id}", VerifyToken(permissionRouter.deactivatePermission, config, dbClient)).Methods("DELETE")

	return router
}

// ---- permissionRouter.createPermission ----
func (rcvr *permissionRouter) createPermission(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		throw(w, err)
	}
	var p root.Permission
	err = json.Unmarshal(body, &p)
	if err != nil {
		throw(w, err)
	}
	p, err := rcvr.permissionService.CreatePermission(p)
	if err == nil {
		rcvr.respond(w, p)
	} else {
		throw(w, err)
	}
}

// ---- permissionRouter.findActivePermissions ----
func (rcvr *permissionRouter) findActivePermission(w http.ResponseWriter, r *http.Request) {
	var p root.Permission
	p.Active = "Yes"
	permissions, err := rcvr.userService.FindPermission(p)
	if err == nil {
		rcvr.respondSlice(w,permissions)
	} else {
		throw(w,err)
	}
}

// ---- permissionRouter.findPermission ----
func (rcvr *permissionRouter) findPermission(w http.ResponseWriter, r *http.Request) {
	var p root.Permission
	vars := mux.Vars(r)
	p.PermissionId = vars["id"]
	permissions, err := rcvr.permissionService.FindPermission(u)
	if err == nil {
		rcvr.respond(w,permissions[0])
	} else {
		throw(w,err)
	}
}

// ---- permissionRouter.updatePermission ----
func (rcvr *permissionRouter) updatePermission(w http.ResponseWriter, r *http.Request) {
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
	var update root.Permission
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
	var filter root.Permission
	vars := mux.Vars(r)
	filter.PermissionId = vars["id"]
	permission, err := rcvr.permissionService.UpdatePermission(filter,update)
	if err == nil {
		rcvr.respond(w,permission)
	} else {
		throw(w,err)
	}
}

// ---- permissionRouter.activatePermission ----
func (rcvr *permissionRouter) activatePermission(w http.ResponseWriter, r *http.Request) {
	var f root.Permission
	vars := mux.Vars(r)
	f.PermissionId = vars["id"]
	var u root.Permission
	u.Active = "Yes"
	permission, err := rcvr.permissionService.UpdatePermission(f,u)
	if err == nil {
		rcvr.respond(w,permission)
	} else {
		throw(w,err)
	}
}

// ---- permissionRouter.deactivatePeremission ----
func (rcvr *permissionRouter) deactivatePermission(w http.ResponseWriter, r *http.Request) {
	var f root.Permission
	vars := mux.Vars(r)
	f.PermissionId = vars["id"]
	var u root.Permission
	u.Active = "No"
	permission, err := rcvr.permissionService.UpdatePermission(f,u)
	if err == nil {
		rcvr.respond(w,permission)
	} else {
		throw(w,err)
	}
}

// ---- permissionRouter.respond ----
func (rcvr *permissionRouter) respond(w http.ResponseWriter, p root.Permission) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		throw(w,err)
	}
}
// ---- permissionRouter.respondSlice ----
func (rcvr *permissionRouter) respondSlice(w http.ResponseWriter, p []root.Permission) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		throw(w,err)
	}
}
