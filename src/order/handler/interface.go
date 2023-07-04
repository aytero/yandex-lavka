package handler

import (
	"context"
	"yandex-team.ru/bstask/model"
)

type Usecase interface {
	GetOrder(ctx context.Context, userID int64) (*model.Order, error)
	GetOrders(ctx context.Context, limit, offset int32) ([]*model.Order, error)
	CreateOrder(ctx context.Context, orders []*model.CreateOrder) ([]*model.Order, error)
	CompleteOrders(ctx context.Context, orders []*model.CompleteOrder) ([]*model.Order, error)
}
