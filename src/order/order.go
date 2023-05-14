package order

import (
    "context"
    "yandex-team.ru/bstask/model"
)

type Repository interface {
    GetById(ctx context.Context, id int64) (*model.OrderDto, error)
    GetOrders(ctx context.Context, limit, offset int32) ([]model.OrderDto, error)
    CreateOrders(ctx context.Context, ords *model.CreateOrderRequest) ([]model.OrderDto, error)
    UpdateOrders(ctx context.Context, ords []model.CompleteOrder) ([]model.OrderDto, error)
}

type Usecase interface {
    GetOrder(ctx context.Context, userID int64) (*model.OrderDto, error)
    GetOrders(ctx context.Context, limit, offset int32) ([]model.OrderDto, error)
    CreateOrder(ctx context.Context, orders *model.CreateOrderRequest) ([]model.OrderDto, error)
    CompleteOrders(ctx context.Context, orders *model.CompleteOrderRequestDto) ([]model.OrderDto, error)
}
