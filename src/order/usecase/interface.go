package usecase

import (
	"context"
	"yandex-team.ru/bstask/model"
)

type Repository interface {
	GetById(ctx context.Context, id int64) (*model.Order, error)
	GetOrders(ctx context.Context, limit, offset int32) ([]*model.Order, error)
	CreateOrders(ctx context.Context, ords []*model.CreateOrder) ([]*model.Order, error)
	UpdateOrders(ctx context.Context, ords []*model.CompleteOrder) ([]*model.Order, error)
}
