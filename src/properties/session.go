package properties

import (
	"github.com/denysetiawan28/go-log/src/server/logging_config"
	"time"
)

type App struct {
	Logger  logging_config.SystemLogger
	Session Session
}

type Session struct {
	RequestTime     time.Time
	ThreadID        string `json:"_app_thread_id"`
	JourneyID       string `json:"_app_journey_id"`
	ServiceName     string `json:"_app_name"`
	ServiceVersion  string `json:"_app_version"`
	IP              string `json:"_IP"`
	ServicePort     int    `json:"_app_port"`
	ReqURI          string `json:"_app_uri"`
	ReqMethod       string `json:"_app_method"`
	SrcIP           string `json:"_src_ip"`
	Header, Request interface{}
	AdditionalData  map[string]interface{} `json:"_app_data,omitempty"`
	ErrorMessage    string
	ResponseCode    string
}

// NewSessionRequest
// Create session request to echo context and set to struct
func NewSessionRequest(logger logging_config.SystemLogger) *App {
	return &App{
		Logger: logger,
		Session: Session{
			RequestTime: time.Now(),
			Header:      map[string]interface{}{},
			Request:     struct{}{},
		},
	}
}
