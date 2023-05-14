package handler

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "net/http"
    "strconv"
    "time"
    "yandex-team.ru/bstask/courier"
    "yandex-team.ru/bstask/model"
    "yandex-team.ru/bstask/server"
)

// todo middleware & validation

type CourierHandler struct {
    uc courier.Usecase
}

func NewCourierHandler(ouc courier.Usecase) *CourierHandler {
    h := &CourierHandler{
        uc: ouc,
    }
    return h
}

func (h *CourierHandler) SetupRoutes(e *echo.Echo) {
    e.GET("/couriers", h.GetCouriers)
    e.POST("/couriers", h.CreateCourier)
    e.GET("/couriers/:courier_id", h.GetCourier)
    e.GET("/couriers/meta-info/:courier_id", h.GetCourierMetaInfo)

}

// GetCouriers -
func (h *CourierHandler) GetCouriers(ctx echo.Context) error {

    // todo check param == ""
    offset, err := strconv.ParseInt(ctx.QueryParam("offset"), 10, 32)
    if err != nil {
        offset = 0
    }
    limit, err := strconv.ParseInt(ctx.QueryParam("limit"), 10, 32)
    if err != nil {
        limit = 1
    }

    resp, err := h.uc.GetCouriers(ctx.Request().Context(), int32(limit), int32(offset))
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    return ctx.JSON(http.StatusOK, resp)
}

// CreateCourier -
func (h *CourierHandler) CreateCourier(ctx echo.Context) error {
    req := model.CreateCourierRequest{}
    err := ctx.Bind(&req)
    if err != nil {
        log.Errorf("CourierHandler - CreateCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    res, err := h.uc.CreateCourier(ctx.Request().Context(), &req)
    if err != nil {
        log.Errorf("CourierHandler - CreateCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    log.Infof("CourierHandler - CreateCouriers: OK")
    return ctx.JSON(http.StatusOK, res)
}

// GetCourier -
func (h *CourierHandler) GetCourier(ctx echo.Context) error {
    // parse and validate
    courierId, err := strconv.ParseInt(ctx.Param("courier_id"), 10, 64)
    if err != nil {
        //log.Println(http.StatusBadRequest)
        log.Errorf("error - %s\n", err)
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }

    // todo context
    entry, err := h.uc.GetCourier(ctx.Request().Context(), courierId)
    if err != nil {
        log.Errorf("CourierHandler - GetCourier: %w", err)
        return ctx.JSON(http.StatusBadRequest, server.BadRequestResponse{})
    }
    if entry == nil {
        log.Error("CourierHandler - GetCourier: 404NotFound")
        return ctx.JSON(http.StatusNotFound, server.NotFoundResponse{})
    }
    log.Infof("CourierHandler - GetCourier: OK")
    return ctx.JSON(http.StatusOK, entry)
}

// GetCourierMetaInfo -
func (h *CourierHandler) GetCourierMetaInfo(ctx echo.Context) error {
    courierId, err := strconv.ParseInt(ctx.Param("courier_id"), 10, 64)
    if err != nil {
        log.Errorf("CourierHandler - GetCourierMetaInfo: %w", err)
        return nil
    }
    startDate, endDate, err := h.parseDates(ctx)
    if err != nil {
        log.Errorf("CourierHandler - GetCourierMetaInfo: %w", err)
        return nil
    }

    res, err := h.uc.GetCourierMetaInfo(ctx.Request().Context(), courierId, startDate, endDate)
    return ctx.JSON(http.StatusOK, res)
}

func (h CourierHandler) parseDates(ctx echo.Context) (time.Time, time.Time, error) {
    startStr := ctx.QueryParam("startDate")
    startDate, err := time.Parse("2006-01-02", startStr)
    if err != nil {
        return time.Time{}, time.Time{}, err
    }
    endStr := ctx.QueryParam("endDate")
    endDate, err := time.Parse("2006-01-02", endStr)
    if err != nil {
        return time.Time{}, time.Time{}, err
    }
    return startDate, endDate, nil
}

// todo task 4
// CouriersAssignments - Список распределенных заказов
//func (c *Container) CouriersAssignments(ctx echo.Context) error {
//	return ctx.JSON(http.StatusOK, models.HelloWorld{
//		Message: "Hello World",
//	})
//}
