package handler

import (
	"context"
	"yandex-team.ru/bstask/order/handler/dto"
)

type Usecase interface {
	GetOrder(ctx context.Context, userID int64) (*dto.OrderDto, error)
	GetOrders(ctx context.Context, limit, offset int32) ([]*dto.OrderDto, error)
	CreateOrder(ctx context.Context, orders *dto.CreateOrderRequest) ([]*dto.OrderDto, error)
	CompleteOrders(ctx context.Context, orders *dto.CompleteOrderRequestDto) ([]*dto.OrderDto, error)
}
