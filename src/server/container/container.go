package container

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	"github.com/denysetiawan28/go-log/src/server/logging_config"
)

type DefaultContainer struct {
	//#register config
	Config *config.DefaultConfig
	Logger logging_config.SystemLogger
}

func IntializeContainer() *DefaultContainer {

	config := config.ConfigApps("./resources/")

	//initialize Logger
	logger := logging_config.SetupLogger(config.Logs)
	//initialize database
	//db := database.InitializeDatabase(config.Database)

	//row := db.Raw("SELECT 1").Row()
	//fmt.Println(row)

	return &DefaultContainer{
		Config: config,
		Logger: logger,
	}
}
