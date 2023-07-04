package dto

import (
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

func ParseFromEntitySlice(entity []*model.Order) []*OrderDto {
	resp := make([]*OrderDto, 0, len(entity))
	for _, e := range entity {
		resp = append(resp, ParseFromEntity(e))
	}
	return resp
}

func ParseFromEntity(m *model.Order) *OrderDto {
	if m == nil {
		return nil
	}
	return &OrderDto{
		OrderId:       m.OrderId,
		Weight:        m.Weight,
		Regions:       m.Regions,
		DeliveryHours: m.DeliveryHours,
		Cost:          m.Cost,
		CompletedTime: m.CompletedTime,
	}
}
