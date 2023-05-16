package logging_config

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
)

type defaultLogger struct {
	writers []io.Writer

	zapLogger *zap.Logger
	level     LogLevel
}

type Field struct {
	Key string
	Val interface{}
}

func NewLogEngine(conf config.Logs) (ILogEngine, error) {
	defLogger := &defaultLogger{
		writers: make([]io.Writer, 0),
	}

	defLogger.zapLogger = NewZapLogger(conf, 0, defLogger.writers...)

	return defLogger, nil
}

func (d *defaultLogger) Debug(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Info(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	zapLogs := []zap.Field{
		//zap.String("logType", LogTypeSYS),
		zap.String("level", "info"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, message, false, fields...)...)
	d.zapLogger.Info("|", zapLogs...)
}

func (d *defaultLogger) Warn(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Error(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Fatal(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Panic(ctx echo.Context, message string, fields ...Field) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultLogger) Close() error {
	//TODO implement me
	panic("implement me")
}

//func NewLogEngine() (*defaultLogger, error)  {
//	defLogger := &defaultLogger{
//		writers: make([]io.Writer, 0),
//	}
//
//	defLogger.zapLogger = logging_config.NewZapLogger(0, defLogger.writers...)
//
//	return defLogger, nil
//}
