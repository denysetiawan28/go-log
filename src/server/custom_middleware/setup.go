package custom_middleware

import (
	"github.com/denysetiawan28/go-log/src/properties"
	"github.com/denysetiawan28/go-log/src/server/container"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func SetupMiddleware(e *echo.Echo, cont *container.DefaultContainer) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		//AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get("request-id")
			journeyId := c.Request().Header.Get("journey-id")

			if len(reqId) == 0 {
				reqId = ""
			}
			session := properties.NewSessionRequest(cont.Logger)
			session.ThreadID = reqId
			session.JourneyID = journeyId
			session.AppName = cont.Config.Apps.Name
			session.AppVersion = cont.Config.Apps.Version
			port, _ := strconv.Atoi(cont.Config.Server.Port)
			session.Port = port
			session.SrcIP = c.RealIP()
			session.URL = c.Request().URL.String()
			session.Method = c.Request().Method
			session.Header = c.Request().Header
			c.Set("App_Session", session)

			return next(c)
		}
	})
}
