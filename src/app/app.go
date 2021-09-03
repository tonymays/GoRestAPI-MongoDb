package main

import (
	"pkg/configuration"
	"pkg/db"
	"pkg/server"
)

// ---- App Struct for the MicroService ----
type App struct {
	Config configuration.Configuration
	Server *server.Server
}

// ---- App.Init ----
func (rcvr *App) Init(env string) error {
	// Step 1: capture api configuration from conf.json
	var config configuration.Configuration
	config, err := configuration.CaptureAPIConfiguration(env)
	if err != nil {
		return err
	}
	rcvr.Config = config

	// Step 2: establish database connections (and there could be more than 1)
//	dataCache := db.NewTestDataCache([]*db.MetricDataModel{})

	// Step 3: Add router services here (and there could be more than 1)
//	metricService := db.NewMetricService(config, dataCache)

	// Step 3: setup our server
	rcvr.Server = server.NewServer(config, dataCache, metricService)

	// Step 4: initialize test data if any
	rcvr.Server.InitTestData()

	// return nil if we are good
	return nil
}

// ---- App.Run ----
func (rcvr *App) Run() {
	// start the server if the train has everything it needs
	rcvr.Server.Start()
}
