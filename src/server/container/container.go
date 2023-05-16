package container

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	Logs "github.com/denysetiawan28/go-log/src/server/config"
	"go.uber.org/zap"
)

type DefaultContainer struct {
	//#register config
	Config *config.DefaultConfig
	Logger *zap.Logger
}

func IntializeContainer() *DefaultContainer {

	config := config.ConfigApps("./resources/")

	//initialize database
	//db := database.InitializeDatabase(config.Database)

	//row := db.Raw("SELECT 1").Row()
	//fmt.Println(row)

	return &DefaultContainer{
		Config: config,
		Logger: Logs.SetupLogger(),
	}
}
