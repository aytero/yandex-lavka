package model

import "time"

// todo domain, model, dto
// todo db tags

type CreateOrderDto struct {
    Weight        float32  `json:"weight" db:"weight"`
    Regions       int32    `json:"regions" db:"regions"`
    DeliveryHours []string `json:"delivery_hours" db:"delivery_hours"`
    Cost          int32    `json:"cost" db:"cost"`
}

type CreateOrderRequest struct {
    Orders []CreateOrderDto `json:"orders"`
}

// todo db:""

type OrderDto struct {
    OrderId       int64    `json:"order_id" db:"order_id"`
    Weight        float32  `json:"weight" db:"weight"`
    Regions       int32    `json:"regions" db:"regions"`
    DeliveryHours []string `json:"delivery_hours" db:"delivery_hours"`
    Cost          int32    `json:"cost" db:"cost"`
    // todo
    CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
}

type GroupOrders struct {
    GroupOrderId int64      `json:"group_order_id"`
    Orders       []OrderDto `json:"orders"`
}

type CompleteOrderRequestDto struct {
    CompleteInfo []CompleteOrder `json:"complete_info"`
}

type CompleteOrder struct {
    CourierId    int64     `json:"courier_id"`
    OrderId      int64     `json:"order_id"`
    CompleteTime time.Time `json:"complete_time"`
}

// todo where does these go?

type OrderAssignResponse struct {
    Date     string                `json:"date"`
    Couriers []CouriersGroupOrders `json:"couriers"`
}

type CouriersGroupOrders struct {
    CourierId int64         `json:"courier_id"`
    Orders    []GroupOrders `json:"orders"`
}
