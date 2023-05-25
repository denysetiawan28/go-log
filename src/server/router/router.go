package router

import (
	"github.com/denysetiawan28/go-log/src/handler"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
)

func InitializeRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "service up")
	})

	server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	g := server.Group("/api/v1")

	g.GET("/test2", func(c echo.Context) error {
		logs := c.Get("zLog").(*zap.Logger)
		//sess := dat.(*properties.Session)
		logs.Info("asdasd")
		return c.JSON(http.StatusOK, "ok")
	})

	g.GET("/test", func(c echo.Context) error {
		//c.Logger().Info(c.Request())
		ctx := c
		ctx.Logger().Info("Logger Info")
		return c.JSON(http.StatusOK, "Hello")
	})
	g.GET("/hello", handler.HelloHandler.HelloWorld)
}
