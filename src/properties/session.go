package properties

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"time"
)

type CustomContext struct {
	echo.Context
}

type Session struct {
	Logger                  *zap.Logger
	RequestTime             time.Time
	ThreadID                string
	JourneyID               string
	AppName, AppVersion, IP string
	Port                    int
	SrcIP, URL, Method      string
	Header, Request         interface{}
	AdditionalData          map[string]interface{}
	ErrorMessage            string
	ResponseCode            string
}

func NewSessionRequest(logger *zap.Logger) *Session {
	return &Session{
		RequestTime: time.Now(),
		Logger:      logger,
		Header:      map[string]interface{}{},
		Request:     struct{}{},
	}
}
