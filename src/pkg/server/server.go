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
		RolePermissionService: rolePermission,
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
	if rcvr.UserService.CountUser() == 0 {
		// step 1: add root user
		fmt.Println("initializing new server")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("   step 1: creating root user...")
		var user root.User
		user.Username = rcvr.Config.RootUserName
		user.Password = rcvr.Config.RootPassword
		user.Firstname = rcvr.Config.Firstname
		user.Lastname = rcvr.Config.Lastname
		user.Address = rcvr.Config.Address
		user.City = rcvr.Config.City
		user.State = rcvr.Config.State
		user.Zip = rcvr.Config.Zip
		user.Country = rcvr.Config.Country
		user.Email = rcvr.Config.Email
		user.Phone = rcvr.Config.Phone
		user.Created = root.GenTimestamp()
		user.Modified = user.Created
		user, err := rcvr.UserService.CreateUser(user)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		// step 2: add admin role
		fmt.Println("   step 2: creating admin role...")
		var role root.Role
		role.Name = "Admin"
		role.Created = user.Created
		role.Modified = role.Created
		role, err = rcvr.RoleService.CreateRole(role)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		// step 3: add default permission sets
		fmt.Println("   step 3: creating system permissions...")
		var permission root.Permission
		permission.Created = user.Created
		permission.Modified = permission.Created

		permission.Tag = "Can Add User"
		_, err = rcvr.PermissionService.CreatePermission(permission)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		permission.Tag = "Can Edit User"
		_, err = rcvr.PermissionService.CreatePermission(permission)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		permission.Tag = "Can Delete User"
		_, err = rcvr.PermissionService.CreatePermission(permission)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		// step 4: add admin role permissions
		fmt.Println("   step 4: assigning permissions to the admin role...")
		var permissions []string
		permissions = append(permissions, "Can Add User")
		permissions = append(permissions, "Can Edit User")
		permissions = append(permissions, "Can Delete User")

		// reseting role to have just the RoleId
		roleId := role.RoleId
		role = root.Role{}
		role.RoleId = roleId

		// add the role permissions
		err = rcvr.RolePermissionService.SetRolePermissions(role, permissions)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}

		// Step 5: assign root user to Admin role
		fmt.Println("   step 5: assigning root user to Admin role...")
		var userRole root.UserRole
		userRole.UserRoleId, _ = root.GenId()
		userRole.UserId = user.UserId
		userRole.RoleId = role.RoleId
		userRole.Active = "Yes"
		userRole.Created = user.Created
		userRole.Modified = userRole.Created
		err = rcvr.UserService.AssignUserRole(userRole)
		if err != nil {
			fmt.Println("\t\tcritical error ", err)
		}
		fmt.Println("------------------------------------------------------------")
		fmt.Println("new server initialization complete")
	}
}