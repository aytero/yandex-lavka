package routes

import (
    "github.com/labstack/echo/v4"
    "yandex-team.ru/bstask/server"
)

// func SetupRoutes(e *echo.Echo, orderHandler *handler.OrderHandler) { // , ordUsecase *order.Usecase, courUsecase *courier.Usecase) {

func SetupRoutes(e *echo.Echo, orderHandler server.Handler) {
    //var handler *echo.Group
    //NewCourierHandler(e, curUC)

    orderHandler.SetupRoutes(e)
    e.GET("/ping", ping)
    //e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    //e.GET("/swagger/*", echoSwagger.WrapHandler)

}
