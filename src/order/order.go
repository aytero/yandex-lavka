package order

import (
    "context"
    "yandex-team.ru/bstask/handler/dto"
    "yandex-team.ru/bstask/model"
)

type Repository interface {
    GetById(ctx context.Context, id int64) (*model.Order, error)
    GetOrders(ctx context.Context, limit, offset int32) ([]*model.Order, error)
    CreateOrders(ctx context.Context, ords []*model.CreateOrder) ([]*model.Order, error)
    UpdateOrders(ctx context.Context, ords []*model.CompleteOrder) ([]*model.Order, error)
}

type Usecase interface {
    GetOrder(ctx context.Context, userID int64) (*dto.OrderDto, error)
    GetOrders(ctx context.Context, limit, offset int32) ([]*dto.OrderDto, error)
    CreateOrder(ctx context.Context, orders *dto.CreateOrderRequest) ([]*dto.OrderDto, error)
    CompleteOrders(ctx context.Context, orders *dto.CompleteOrderRequestDto) ([]*dto.OrderDto, error)
}
