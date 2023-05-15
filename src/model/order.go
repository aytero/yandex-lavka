package model

import "time"

type Order struct {
    OrderId       int64      `json:"order_id" db:"order_id"`
    Weight        float32    `json:"weight" db:"weight"`
    Regions       int32      `json:"regions" db:"regions"`
    DeliveryHours []string   `json:"delivery_hours" db:"delivery_hours"`
    Cost          int32      `json:"cost" db:"cost"`
    CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
}

type CreateOrder struct {
    Weight        float32  `json:"weight" db:"weight"`
    Regions       int32    `json:"regions" db:"regions"`
    DeliveryHours []string `json:"delivery_hours" db:"delivery_hours"`
    Cost          int32    `json:"cost" db:"cost"`
}

type CompleteOrder struct {
    CourierId    int64     `json:"courier_id" db:"courier_id"`
    OrderId      int64     `json:"order_id" db:"order_id"`
    CompleteTime time.Time `json:"complete_time" db:"complete_time"`
}
