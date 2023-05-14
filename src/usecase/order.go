package usecase

import (
    "context"
    "yandex-team.ru/bstask/model"
    "yandex-team.ru/bstask/order"
)

type OrderUsecase struct {
    orderRepo order.Repository
}

func NewOrderUsecase(repo order.Repository) *OrderUsecase {
    return &OrderUsecase{
        orderRepo: repo,
    }
}

func (uc *OrderUsecase) GetOrder(ctx context.Context, orderId int64) (*model.OrderDto, error) {
    entry, err := uc.orderRepo.GetById(ctx, orderId) // uc.BalanceRepo.GetById ; getUserByID
    if err != nil {
        //return nil, fmt.Errorf("OrderRepository - CreateOrders: %w", err)
        return nil, err
    }
    return entry, nil
}

func (uc *OrderUsecase) GetOrders(ctx context.Context, limit, offset int32) ([]model.OrderDto, error) {
    entry, err := uc.orderRepo.GetOrders(ctx, limit, offset)
    if err != nil {
        // todo logging
        return nil, err
    }
    return entry, nil
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, orders *model.CreateOrderRequest) ([]model.OrderDto, error) {
    entry, err := uc.orderRepo.CreateOrders(ctx, orders)
    if err != nil {
        return nil, err
    }
    return entry, nil
}

func (uc *OrderUsecase) CompleteOrders(ctx context.Context, orders *model.CompleteOrderRequestDto) ([]model.OrderDto, error) {
    entry, err := uc.orderRepo.UpdateOrders(ctx, orders.CompleteInfo)
    if err != nil {
        return nil, err
    }
    return entry, nil
}
