package dto

import (
	"time"
	"yandex-team.ru/bstask/model"
)

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

func (o CompleteOrder) MapToModel() model.CompleteOrder {
	return model.CompleteOrder{
		CourierId:    o.CourierId,
		OrderId:      o.OrderId,
		CompleteTime: o.CompleteTime,
	}
}

func (or CompleteOrderRequestDto) MapToModel() []*model.CompleteOrder {
	req := make([]*model.CompleteOrder, 0, len(or.CompleteInfo))
	for _, o := range or.CompleteInfo {
		r := o.MapToModel()
		req = append(req, &r)
	}
	return req
}
