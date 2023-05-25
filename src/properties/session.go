package properties

import (
	log_watcher "github.com/denysetiawan28/log-watcher"
)

type App struct {
	Logger log_watcher.SystemLogger
}

// NewSessionRequest
// Create session request to echo context and set to struct
func NewSessionRequest(logger log_watcher.SystemLogger) *App {
	return &App{
		Logger: logger,
	}
}
