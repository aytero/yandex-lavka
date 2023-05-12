package handler

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "log"
    "net/http"
    "strconv"
    "yandex-team.ru/bstask/model"
    "yandex-team.ru/bstask/server"
    "yandex-team.ru/bstask/usecase"
)

type OrderHandler struct {
    uc *usecase.OrderUsecase
}

// todo middleware & validation

func NewOrderHandler(e *echo.Echo, ouc *usecase.OrderUsecase) *OrderHandler {
    h := &OrderHandler{
        uc: ouc,
    }

    // todo diff func smth like settings / setup routes
    e.GET("/orders", h.GetOrders)
    e.POST("/orders", h.CreateOrder)
    e.POST("/orders/complete", h.CompleteOrder)
    //e.POST("/orders/assign", h.OrdersAssign)
    e.GET("/orders/:order_id", h.GetOrder)

    return h
}

// GetOrders -
func (h *OrderHandler) GetOrders(ctx echo.Context) error {

    // todo check param == ""
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
        ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    //return h.JSON(http.StatusOK, newOrdersResponse(h.uc.OrderRepo, userIDFromToken(h), orders, count))
    return ctx.JSON(http.StatusOK, orders)
}

// CreateOrder -
func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
    req := model.CreateOrderRequest{}
    err := ctx.Bind(&req)
    if err != nil {
        log.Println(fmt.Errorf("OrderHandler - CreateOrders: %w", err))
        ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    res, err := h.uc.CreateOrder(ctx.Request().Context(), &req)
    if err != nil {
        log.Println(fmt.Errorf("OrderHandler - CreateOrders: %w", err))
        ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    log.Println("OrderHandler - CreateOrders: OK")
    return ctx.JSON(http.StatusOK, res)
}

// CompleteOrder -
func (h *OrderHandler) CompleteOrder(ctx echo.Context) error {
    req := &model.CompleteOrderRequestDto{}
    err := ctx.Bind(&req)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }

    res, err := h.uc.CompleteOrders(ctx.Request().Context(), req)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    return ctx.JSON(http.StatusOK, res)
}

// GetOrder -
func (h *OrderHandler) GetOrder(ctx echo.Context) error {
    // parse and validate
    orderId, err := strconv.ParseInt(ctx.Param("order_id"), 10, 64)
    if err != nil {
        //log.Println(http.StatusBadRequest)
        log.Printf("error - %s\n", err)
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }

    // todo context
    entry, err := h.uc.GetOrder(ctx.Request().Context(), orderId)
    if err != nil {

        log.Println(fmt.Errorf("OrderHandler - GetOrder: %w", err))
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    if entry == nil {
        log.Println("OrderHandler - GetOrder: 404NotFound")
        return ctx.JSON(http.StatusNotFound, server.NotFoundResponse{})
    }
    log.Println("OrderHandler - GetOrder: OK")
    return ctx.JSON(http.StatusOK, entry)
}

// todo task 4
// OrdersAssign - Распределение заказов по курьерам
//func (c *OrderHandler) OrdersAssign(ctx echo.Context) error {
//    if err != nil {
//        ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
//    }
//    return ctx.JSON(http.StatusCreated, model.OrderAssignResponse{})
//}
