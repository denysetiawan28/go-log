package handler

import (
	"github.com/denysetiawan28/go-log/src/constanta/constant"
	"github.com/denysetiawan28/go-log/src/server/container"
	"github.com/labstack/echo/v4"
	"net/http"
)

type helloHandler struct {
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h helloHandler) HelloWorld(ctx echo.Context) (err error) {
	sess := ctx.Get(constant.AppLoggerID).(*container.AppLogger)
	//ctxi := ctx.Request().Context().
	sess.Logger.Info(sess.LogContext, "Hello From Logger")
	//sess.Logger.Info(nil, "testing")
	return ctx.JSON(http.StatusOK, "ok")
}
