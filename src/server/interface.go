package server

import "github.com/labstack/echo/v4"

type Handler interface {
	SetupRoutes(e *echo.Echo)
}
