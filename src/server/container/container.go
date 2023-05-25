package container

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	log_watcher "github.com/denysetiawan28/log-watcher"
	"golang.org/x/net/context"
)

type DefaultContainer struct {
	//#register config
	Config *config.DefaultConfig
}

type AppLogger struct {
	LogContext context.Context
	Logger     log_watcher.SystemLogger
}

func IntializeContainer() (*DefaultContainer, *AppLogger) {

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
	defContainer := &DefaultContainer{
		Config: config,
	}

	appLogger := &AppLogger{
		Logger: logger,
	}

	return defContainer, appLogger
}
