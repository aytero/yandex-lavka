package server

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "github.com/rs/zerolog"
    "net/http"
    "os"
)

func (s *Server) SetupMiddleware(e *echo.Echo) {

    logger := zerolog.New(os.Stdout)
    switch s.cfg.LogLevel {
    case "off":
        s.server.Logger.SetLevel(log.OFF)
    case "error":
        s.server.Logger.SetLevel(log.ERROR)
    case "warn":
        s.server.Logger.SetLevel(log.WARN)
    case "info":
        s.server.Logger.SetLevel(log.INFO)
    case "debug":
        s.server.Logger.SetLevel(log.ERROR)
    default:
        s.server.Logger.SetLevel(log.WARN)
    }

    e.Use(middleware.Recover())
    e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
        LogURI:    true,
        LogStatus: true,
        LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
            logger.Info().
                Str("URI", v.URI).
                Int("status", v.Status).
                Msg("request")

            return nil
        },
    }))
    e.Use(middleware.Logger())

    e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
        Skipper: middleware.DefaultSkipper,
        Store:   middleware.NewRateLimiterMemoryStore(10),
        DenyHandler: func(context echo.Context, identifier string, err error) error {
            return context.JSON(http.StatusTooManyRequests, nil)
        },
    }))

}
