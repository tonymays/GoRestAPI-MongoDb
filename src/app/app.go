package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pkg/configuration"
	"pkg/server"
	"pkg/services"
	"time"
)

// ---- App Struct ----
type App struct {
	Server	*server.Server
}

// ---- App.Init ----
func (rcvr *App) Init(e string) error {
	// Step 1: Load Configuration Settings and set the environment
	config, err := configuration.Init(e)
	if err != nil {
		return err
	}
	config.Environment = e

	// Step 2: Connect to the targeted Mongo Database
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	err = dbClient.Connect(ctx)
	if err != nil {
		return err
	}

	// Step 3: Setup App Services
	roleService := services.NewRoleService(config, dbClient)
	permissionService := services.NewPermissionService(config, dbClient)
	rolePermissionService := services.NewRolePermissionService(config, dbClient, roleService, permissionService)
	userService := services.NewUserService(config, dbClient, roleService, rolePermissionService)
	authService := services.NewAuthService(config, dbClient, userService)

	// Step 4: Setup the App Server
	rcvr.Server = server.NewServer(
		config,
		dbClient,
		authService,
		userService,
		roleService,
		permissionService,
		rolePermissionService,
	)

	// Step 5: Perform any Server Startup functions
	rcvr.Server.Init()

	// Step 6: Return nil letting the main go routine that dependency injection
	//         of the App and Server structures are now complete and the
	//         microservice is ready to run
	return nil
}

// ---- App.Run ----
func (rcvr *App) Run() {
	// start the server
	rcvr.Server.Start()
}