package logging_config

import (
	"github.com/denysetiawan28/go-log/src/constanta/constant"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ctxKeyLogger struct{}

var key ctxKeyLogger

func formatLogs(ctx echo.Context, msg string, mask bool, fields ...Field) (logRecord []zap.Field) {
	sess := extractSession(ctx)

	// add global value from context that must be exist on all logs!
	logRecord = append(logRecord, zap.String("message", msg))

	logRecord = append(logRecord, zap.String("app_name", sess.ServiceName))
	logRecord = append(logRecord, zap.String("app_tag", sess.Tag))
	logRecord = append(logRecord, zap.String("app_version", sess.ServiceVersion))
	logRecord = append(logRecord, zap.Int("app_port", sess.ServicePort))
	logRecord = append(logRecord, zap.String("app_thread_id", sess.ThreadID))
	logRecord = append(logRecord, zap.String("app_journey_id", sess.JourneyID))
	logRecord = append(logRecord, zap.String("app_req_ip", sess.SrcIP))
	logRecord = append(logRecord, zap.String("app_method", sess.ReqMethod))
	logRecord = append(logRecord, zap.String("app_uri", sess.ReqURI))

	if sess.Request != nil {
		logRecord = append(logRecord, zap.Any("app_req_body", sess.Request))
	} else {
		logRecord = append(logRecord, zap.Any("app_req_body", nil))
	}

	// add additional data that available across all log, such as user_id
	if sess.AdditionalData != nil {
		logRecord = append(logRecord, zap.Any("app_additional_data", sess.AdditionalData))
	} else {
		logRecord = append(logRecord, zap.Any("app_additional_data", nil))
	}

	return
}

func extractSession(ctx echo.Context) Context {
	if ctx == nil {
		return Context{}
	}

	val, err := ctx.Get(constant.AppSessionID).(Context)

	if !err {
		return Context{}
	}

	return val
}
