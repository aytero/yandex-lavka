package dto

import (
    "errors"
    "regexp"
    "time"
    "yandex-team.ru/bstask/model"
)

type OrderDto struct {
    OrderId       int64      `json:"order_id"`
    Weight        float32    `json:"weight"`
    Regions       int32      `json:"regions"`
    DeliveryHours []string   `json:"delivery_hours"`
    Cost          int32      `json:"cost"`
    CompletedTime *time.Time `json:"completed_time,omitempty"`
}

type CreateOrderDto struct {
    Weight        float32  `json:"weight"`
    Regions       int32    `json:"regions"`
    DeliveryHours []string `json:"delivery_hours"`
    Cost          int32    `json:"cost"`
}

func (o CreateOrderDto) MapToModel() model.CreateOrder {
    return model.CreateOrder{
        Weight:        o.Weight,
        Regions:       o.Regions,
        DeliveryHours: o.DeliveryHours,
        Cost:          o.Cost,
    }
}

type CreateOrderRequest struct {
    Orders []CreateOrderDto `json:"orders"`
}

func (o *CreateOrderRequest) Validate() error {
    timeRegex := regexp.MustCompile(`^([0-1][0-9]|2[0-3]):[0-5][0-9]-([0-1][0-9]|2[0-3]):[0-5][0-9]$`)
    for _, order := range o.Orders {
        //if len(order.DeliveryHours) == 0 || order.Cost < 0 || order.Regions < 0 || order.Weight < 0 {
        if len(order.DeliveryHours) == 0 {
            return errors.New("json validation error")
        }
        for _, interval := range order.DeliveryHours {
            if err := timeRegex.MatchString(interval); !err {
                return errors.New("json validation error")
            }
        }
    }
    return nil
}

type CompleteOrderRequestDto struct {
    CompleteInfo []CompleteOrder `json:"complete_info"`
}

type CompleteOrder struct {
    CourierId    int64     `json:"courier_id"`
    OrderId      int64     `json:"order_id"`
    CompleteTime time.Time `json:"complete_time"`
}

func (o CompleteOrder) MapToModel() model.CompleteOrder {
    return model.CompleteOrder{
        CourierId:    o.CourierId,
        OrderId:      o.OrderId,
        CompleteTime: o.CompleteTime,
    }
}

type GroupOrders struct {
    GroupOrderId int64      `json:"group_order_id"`
    Orders       []OrderDto `json:"orders"`
}
