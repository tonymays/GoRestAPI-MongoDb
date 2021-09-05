package main

import (
	"pkg/data"
	"pkg/server"
)

type App struct {
	Server	*server.Server
}

func (rcvr *App) Init(e string) error {
	var data data.Data
	data.Init(e)
	rcvr.Server = server.NewServer(data)
	rcvr.Server.Init()
	return nil
}

func (rcvr *App) Run() {
	rcvr.Server.Start()
}