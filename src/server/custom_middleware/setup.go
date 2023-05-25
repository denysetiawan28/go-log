package custom_middleware

import (
	"github.com/denysetiawan28/go-log/src/constanta/constant"
	"github.com/denysetiawan28/go-log/src/server/container"
	log_watcher "github.com/denysetiawan28/log-watcher"
	echoprometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"golang.org/x/net/context"
	"io/ioutil"
	"strconv"
	"time"
)

func SetupMiddleware(e *echo.Echo, cont *container.DefaultContainer, appLogger *container.AppLogger) {
	// Set CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		//AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
	}))

	e.Use(echoprometheus.MetricsMiddleware())

	e.Use(middleware.RequestID())

	// This middleware has many function
	// 1. Set every request to application and set the request echo context
	// 2. Set logger to echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get("X-Request-Id")
			journeyId := c.Request().Header.Get("X-Journey-Id")

			if len(reqId) == 0 {
				reqId = random.String(32)
				c.Request().Header.Set("X-Request-Id", reqId)
			}

			//app := properties.NewSessionRequest(cont.Logger)

			port, _ := strconv.Atoi(cont.Config.Server.Port)
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				body = nil
			}

			dt := log_watcher.Context{
				RequestTime:    time.Time{},
				RequestID:      reqId,
				JourneyID:      journeyId,
				ServiceName:    cont.Config.Apps.Name,
				ServiceVersion: cont.Config.Apps.Version,
				ServicePort:    port,
				ReqURI:         c.Request().URL.String(),
				Tag:            cont.Config.Apps.Tag,
				ReqMethod:      c.Request().Method,
				SrcIP:          c.RealIP(),
				Header:         c.Request().Header,
				Request:        string(body),
			}

			//set log information to golang context
			ctx := log_watcher.SetContext(context.Background(), "", dt)
			//c.SetRequest(c.Request().Clone(ctx))
			// Set Log Info To Context
			appLogger.LogContext = ctx
			c.Set(constant.AppLoggerID, appLogger)
			//c.Set(constant.AppSessionID, dt)

			return next(c)
		}
	})
}
