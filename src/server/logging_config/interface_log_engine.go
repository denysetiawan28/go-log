package logging_config

import (
	"github.com/labstack/echo/v4"
)

type ILogEngine interface {
	Debug(ctx echo.Context, message string, fields ...Field)
	Info(ctx echo.Context, message string, fields ...Field)
	Warn(ctx echo.Context, message string, fields ...Field)
	Error(ctx echo.Context, message string, fields ...Field)
	Fatal(ctx echo.Context, message string, fields ...Field)
	Panic(ctx echo.Context, message string, fields ...Field)
	Close() error
}
