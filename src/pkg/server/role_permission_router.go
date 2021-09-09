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
type rolePermissionRouter struct {
	config					configuration.Configuration
	dbClient				*mongo.Client
	rolePermissionService	root.RolePermissionService
}

// ---- NewUserRouter ----
func NewRolePermissionRouter(router *mux.Router, config configuration.Configuration, dbClient *mongo.Client, rolePermissionService root.RolePermissionService) *mux.Router {
	rolePermissionRouter :=  rolePermissionRouter{config, dbClient, rolePermissionService}

	router.HandleFunc("/role/{id}/permissions", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/role/{id}/permissions", VerifyToken(rolePermissionRouter.setRolePermissions, config, dbClient)).Methods("POST")
	router.HandleFunc("/role/{id}/permissions", VerifyToken(rolePermissionRouter.findRolePermissions, config, dbClient)).Methods("GET")

	return router
}

// ---- rolePermissionRouter.setRolePermissions ----
func (rcvr *rolePermissionRouter) setRolePermissions(w http.ResponseWriter, r *http.Request) {
	// get the specified role id given
	vars := mux.Vars(r)
	var role root.Role
	role.RoleId = vars["id"]

	// grab the permission set
	var permissionTags []string
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
	err = json.Unmarshal(body, &permissionTags)
	if err != nil {
		throw(w,err)
		return
	}
	err = rcvr.rolePermissionService.SetRolePermissions(role, permissionTags)
	if err == nil {
		w = SetResponseHeaders(w, "", "")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(permissionTags)
	} else {
		throw(w, err)
	}
}

// ---- rolePermissionRouter.findRolePermissions ----
func (rcvr *rolePermissionRouter) findRolePermissions(w http.ResponseWriter, r *http.Request) {
	// get the specified role id given
	vars := mux.Vars(r)
	var role root.Role
	role.RoleId = vars["id"]
	rolePermissions, err := rcvr.rolePermissionService.FindRolePermission(role)
	if err == nil {
		w = SetResponseHeaders(w, "", "")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rolePermissions)
	} else {
		throw(w, err)
	}
}