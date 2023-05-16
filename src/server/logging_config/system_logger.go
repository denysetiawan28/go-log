package logging_config

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	"time"
)

type SystemLogger struct {
	ILogEngine
}

type Context struct {
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
	Tag             string `json:"_app_tag"`
}

func SetupLogger(conf config.Logs) SystemLogger {
	sysLog, _ := NewLogEngine(conf)

	return SystemLogger{sysLog}
}
