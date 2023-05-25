package container

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	log_watcher "github.com/denysetiawan28/log-watcher"
)

type DefaultContainer struct {
	//#register config
	Config *config.DefaultConfig
	Logger log_watcher.SystemLogger
}

func IntializeContainer() *DefaultContainer {

	config := config.ConfigApps("./resources/")

	//initialize Logger
	logger := log_watcher.SetupLogger(log_watcher.LogConfig{
		Stdout:           config.Logs.Stdout,
		File:             config.Logs.File,
		Path:             config.Logs.Path,
		MaximumLogSize:   config.Logs.MaximumLogSize,
		MaximumLogAge:    config.Logs.MaximumLogAge,
		MaximumLogBackup: config.Logs.MaximumLogBackup,
	})
	//initialize database
	//db := database.InitializeDatabase(config.Database)

	//row := db.Raw("SELECT 1").Row()
	//fmt.Println(row)

	return &DefaultContainer{
		Config: config,
		Logger: logger,
	}
}
