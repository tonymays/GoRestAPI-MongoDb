package main

import (
	"pkg"
	"pkg/configuration"
	"pkg/server"
	"pkg/services"
)

type App struct {
	Server	*server.Server
}

func (rcvr *App) Init(e string) error {
	config, err := configuration.Init(e)
	if err != nil {
		return err
	}

	dbClient, err := root.Connect(config.MongoUri)
	if err != nil {
		return err
	}

	dbService := services.NewDbService(config, dbClient)
	authService := services.NewAuthService(config, dbClient, dbService)
	userService := services.NewUserService(config, dbClient, dbService)

	rcvr.Server = server.NewServer(config, dbClient, dbService, authService, userService)

	rcvr.Server.Init()
	return nil
}

func (rcvr *App) Run() {
	rcvr.Server.Start()
}