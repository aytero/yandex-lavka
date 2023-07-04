package dto

import (
	"errors"
	"regexp"
	"yandex-team.ru/bstask/model"
)

type CreateOrderDto struct {
	Weight        float32  `json:"weight"`
	Regions       int32    `json:"regions"`
	DeliveryHours []string `json:"delivery_hours"`
	Cost          int32    `json:"cost"`
}

type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders"`
}

func (o CreateOrderDto) MapToModel() model.CreateOrder {
	return model.CreateOrder{
		Weight:        o.Weight,
		Regions:       o.Regions,
		DeliveryHours: o.DeliveryHours,
		Cost:          o.Cost,
	}
}

func (o CreateOrderRequest) MapToModel() []*model.CreateOrder {
	req := make([]*model.CreateOrder, 0, len(o.Orders))
	for _, ord := range o.Orders {
		r := ord.MapToModel()
		req = append(req, &r)
	}
	return req
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
