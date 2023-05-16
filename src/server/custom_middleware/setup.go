package custom_middleware

import (
	"github.com/denysetiawan28/go-log/src/constanta/constant"
	"github.com/denysetiawan28/go-log/src/properties"
	"github.com/denysetiawan28/go-log/src/server/container"
	"github.com/denysetiawan28/go-log/src/server/logging_config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"strconv"
	"time"
)

func SetupMiddleware(e *echo.Echo, cont *container.DefaultContainer) {
	// Set CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		//AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
	}))

	// This middleware has many function
	// 1. Set every request to application and set the request echo context
	// 2. Set logger to echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get("request-id")
			journeyId := c.Request().Header.Get("journey-id")

			if len(reqId) == 0 {
				reqId = ""
			}
			app := properties.NewSessionRequest(cont.Logger)

			port, _ := strconv.Atoi(cont.Config.Server.Port)
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				body = nil
			}

			dt := logging_config.Context{
				RequestTime:    time.Time{},
				ThreadID:       reqId,
				JourneyID:      journeyId,
				ServiceName:    cont.Config.Apps.Name,
				ServiceVersion: cont.Config.Apps.Version,
				IP:             "",
				ServicePort:    port,
				ReqURI:         c.Request().URL.String(),
				Tag:            cont.Config.Apps.Tag,
				ReqMethod:      c.Request().Method,
				SrcIP:          c.RealIP(),
				Header:         c.Request().Header,
				Request:        body,
				AdditionalData: nil,
				ErrorMessage:   "",
				ResponseCode:   "",
			}

			c.Set(constant.AppLoggerID, app)
			c.Set(constant.AppSessionID, dt)

			return next(c)
		}
	})
}
