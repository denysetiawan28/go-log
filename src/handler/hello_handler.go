package handler

import (
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
	dat := ctx.Get("App_Session")
	sess := dat.(*properties.Session)
	print(sess)
	sess.Logger.Info("testing1")
	sess.Logger.Info("testing")
	return ctx.JSON(http.StatusOK, "ok")
}
