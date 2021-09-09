package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"pkg"
	"pkg/configuration"
)

// ---- Server Struct ----
type Server struct {
	Router					*mux.Router
	Config					configuration.Configuration
	DbClient				*mongo.Client
	AuthService				root.AuthService
	UserService				root.UserService
	RoleService				root.RoleService
	PermissionService		root.PermissionService
	RolePermissionService	root.RolePermissionService
}

// --- NewServer ----
func NewServer(config configuration.Configuration, dbClient *mongo.Client, auth root.AuthService, user root.UserService, role root.RoleService, permission root.PermissionService, rolePermission root.RolePermissionService) *Server {
	// establish routers
	router := mux.NewRouter().StrictSlash(true)
	router = NewAuthRouter(router, config, dbClient, auth)
	router = NewUserRouter(router, config, dbClient, user)
	router = NewRoleRouter(router, config, dbClient, role)
	router = NewPermissionRouter(router, config, dbClient, permission)
	router = NewRolePermissionRouter(router, config, dbClient, rolePermission)

	// setup Server struct
	s := Server{
		Router: router,
		Config: config,
		DbClient: dbClient,
		AuthService: auth,
		UserService: user,
		RoleService: role,
		PermissionService: permission,
	}

	// return the Server struct
	return &s
}

// ---- Server.Start ----
func (rcvr *Server) Start() {
	// if https is on ...
	if rcvr.Config.HTTPS == "on" {
		fmt.Println("Listening on port 8443")
		http.ListenAndServeTLS(":8443", rcvr.Config.Cert, rcvr.Config.Key, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	// otherwise...
	} else {
		fmt.Println("Listening on port", rcvr.Config.ServerListenPort)
		http.ListenAndServe(rcvr.Config.ServerListenPort, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	}
}

// ---- Server.Init ----
func (rcvr *Server) Init() {
	// Check user count if zero add root user

	// Step 1: Add Default Permission Sets

	// Step 2: Add Admin Role

	// Step 3: Add Admin Role Permissions

	// Step 4: Add default root user here

	// Step 5: Assign root user to Admin Role
}