package usecase

import (
	"context"
	"yandex-team.ru/bstask/model"
)

type OrderUsecase struct {
	orderRepo Repository
}

func NewOrderUsecase(repo Repository) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: repo,
	}
}

func (uc *OrderUsecase) GetOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	entry, err := uc.orderRepo.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return entry, nil

}

func (uc *OrderUsecase) GetOrders(ctx context.Context, limit, offset int32) ([]*model.Order, error) {
	entry, err := uc.orderRepo.GetOrders(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, orders []*model.CreateOrder) ([]*model.Order, error) {

	entry, err := uc.orderRepo.CreateOrders(ctx, orders)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (uc *OrderUsecase) CompleteOrders(ctx context.Context, orders []*model.CompleteOrder) ([]*model.Order, error) {
	entry, err := uc.orderRepo.UpdateOrders(ctx, orders)
	if err != nil {
		return nil, err
	}

	return entry, nil
}
