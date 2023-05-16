package handler

import (
	"github.com/denysetiawan28/go-log/src/constanta/constant"
	"github.com/denysetiawan28/go-log/src/properties"
	"github.com/labstack/echo/v4"
	"net/http"
)

type helloHandler struct {
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h helloHandler) HelloWorld(ctx echo.Context) (err error) {
	dat := ctx.Get(constant.AppLoggerID)
	sess := dat.(*properties.App)
	print(sess)
	sess.Logger.Info(ctx, "testing1")
	//sess.Logger.Info(nil, "testing")
	return ctx.JSON(http.StatusOK, "ok")
}
