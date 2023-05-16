package config

import "time"

type DefaultConfig struct {
	Apps     Apps       `mapstructure:"apps"`
	Server   Server     `mapstructure:"server"`
	Logs     Logs       `mapstructure:"logs"`
	Database Datasource `mapstructure:"database"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Tag     string `mapstructure:"tag"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Logs struct {
	Stdout           bool   `mapstructure:"stdout"`
	File             bool   `mapstructure:"file"`
	MaximumLogSize   int    `mapstructure:"maximumLogSize"`
	MaximumLogBackup int    `mapstructure:"maximumLogBackup"`
	MaximumLogAge    int    `mapstructure:"maximumLogAge"`
	Path             string `mapstructure:"path"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxOpenConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}
