package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	dtoOrder "yandex-team.ru/bstask/order/handler/dto"
)

type OrderHandler struct {
	uc Usecase
}

func NewOrderHandler(ouc Usecase) *OrderHandler {
	h := &OrderHandler{
		uc: ouc,
	}
	return h
}

func (h *OrderHandler) SetupRoutes(e *echo.Echo) {
	e.GET("/orders", h.GetOrders)
	e.POST("/orders", h.CreateOrder)
	e.POST("/orders/complete", h.CompleteOrder)
	//e.POST("/orders/assign", h.OrdersAssign)
	e.GET("/orders/:order_id", h.GetOrder)
}

// GetOrder -
func (h *OrderHandler) GetOrder(ctx echo.Context) error {
	orderId, err := strconv.ParseInt(ctx.Param("order_id"), 10, 64)
	if err != nil {
		log.Infof("error - %v\n", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}

	entry, err := h.uc.GetOrder(ctx.Request().Context(), orderId)
	if err != nil {
		log.Infof("OrderHandler - GetOrder: %w", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}
	if entry == nil {
		log.Info("OrderHandler - GetOrder: 404NotFound")
		return ctx.JSON(http.StatusNotFound, dtoOrder.NotFoundResponse{})
	}
	log.Info("OrderHandler - GetOrder: OK")
	return ctx.JSON(http.StatusOK, entry)
}

// GetOrders -
func (h *OrderHandler) GetOrders(ctx echo.Context) error {

	offset, err := strconv.ParseInt(ctx.QueryParam("offset"), 10, 32)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(ctx.QueryParam("limit"), 10, 32)
	if err != nil {
		limit = 1
	}
	orders, err := h.uc.GetOrders(ctx.Request().Context(), int32(limit), int32(offset))
	if err != nil {
		log.Infof("OrderHandler - GetOrders: BadRequest: %v", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}
	log.Info("OrderHandler - GetOrders: OK")
	return ctx.JSON(http.StatusOK, orders)
}

// CreateOrder -
func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	req := dtoOrder.CreateOrderRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		log.Infof("OrderHandler: %v", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}
	res, err := h.uc.CreateOrder(ctx.Request().Context(), &req)
	if err != nil {
		log.Infof("OrderHandler - CreateOrders: %v", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}
	return ctx.JSON(http.StatusOK, res)
}

// CompleteOrder -
func (h *OrderHandler) CompleteOrder(ctx echo.Context) error {
	req := &dtoOrder.CompleteOrderRequestDto{}
	err := ctx.Bind(&req)
	if err != nil {
		log.Infof("OrderHandler - CompleteOrder: %w", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}

	res, err := h.uc.CompleteOrders(ctx.Request().Context(), req)
	if err != nil {
		log.Infof("OrderHandler - CompleteOrder: %v", err)
		return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
	}
	log.Info("OrderHandler - CompleteOrder: OK")
	return ctx.JSON(http.StatusOK, res)
}

// todo task 4
// OrdersAssign - Распределение заказов по курьерам
//func (c *OrderHandler) OrdersAssign(ctx echo.Context) error {
//    if err != nil {
//        return ctx.JSON(http.StatusBadRequest, dtoOrder.BadRequestResponse{})
//    }
//    return ctx.JSON(http.StatusCreated, dtoOrder.OrderAssignResponse{})
//}
