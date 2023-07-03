package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"time"
	dtoCourier "yandex-team.ru/bstask/courier/handler/dto"
)

type CourierHandler struct {
    uc Usecase
}

func NewCourierHandler(ouc Usecase) *CourierHandler {
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
        log.Infof("CourierHandler - GetCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }
    return ctx.JSON(http.StatusOK, resp)
}

// CreateCourier -
func (h *CourierHandler) CreateCourier(ctx echo.Context) error {
    req := dtoCourier.CreateCourierRequest{}
    err := ctx.Bind(&req)
    if err != nil {
        log.Infof("CourierHandler - CreateCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }
    if err := req.Validate(); err != nil {
        log.Infof("CourierHandler - CreateCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }
    res, err := h.uc.CreateCourier(ctx.Request().Context(), &req)
    if err != nil {
        log.Infof("CourierHandler - CreateCouriers: %w", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }
    log.Info("CourierHandler - CreateCouriers: OK")
    return ctx.JSON(http.StatusOK, res)
}

// GetCourier -
func (h *CourierHandler) GetCourier(ctx echo.Context) error {
    // parse and validate
    courierId, err := strconv.ParseInt(ctx.Param("courier_id"), 10, 64)
    if err != nil {
        log.Infof("error - %s\n", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }

    // todo context
    entry, err := h.uc.GetCourier(ctx.Request().Context(), courierId)
    if err != nil {
        log.Infof("CourierHandler - GetCourier: %w", err)
        return ctx.JSON(http.StatusBadRequest, dtoCourier.BadRequestResponse{})
    }
    if entry == nil {
        log.Infof("CourierHandler - GetCourier: 404NotFound")
        return ctx.JSON(http.StatusNotFound, dtoCourier.NotFoundResponse{})
    }
    log.Info("CourierHandler - GetCourier: OK")
    return ctx.JSON(http.StatusOK, entry)
}

// GetCourierMetaInfo -
func (h *CourierHandler) GetCourierMetaInfo(ctx echo.Context) error {
    courierId, err := strconv.ParseInt(ctx.Param("courier_id"), 10, 64)
    if err != nil {
        log.Infof("CourierHandler - GetCourierMetaInfo: %w", err)
        return nil
    }
    startDate, endDate, err := h.parseDates(ctx)
    if err != nil {
        log.Infof("CourierHandler - GetCourierMetaInfo: %w", err)
        return nil
    }

    res, err := h.uc.GetCourierMetaInfo(ctx.Request().Context(), courierId, startDate, endDate)
    return ctx.JSON(http.StatusOK, res)
}

// CouriersAssignments - Список распределенных заказов
func (h *CourierHandler) CouriersAssignments(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, dtoCourier.OrderAssignResponse{})
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
