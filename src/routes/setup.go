package routes

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/handler"
)

/*
func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", ping)
}
*/

func SetupRoutes(e *echo.Echo, orderHandler *handler.OrderHandler) { // , ordUsecase *order.Usecase, courUsecase *courier.Usecase) {
	//var handler *echo.Group
	//NewOrderHandler(handler)

	e.GET("/ping", ping)
	//e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	//NewCourierHandler(e, curUC)
	//
	e.GET("/orders", orderHandler.GetOrders)
	e.POST("/orders", orderHandler.CreateOrder)
	e.POST("/orders/complete", orderHandler.CompleteOrder)
	//e.POST("/orders/assign", orderHandler.OrdersAssign)
	e.GET("/orders/:order_id", orderHandler.GetOrder)
}
