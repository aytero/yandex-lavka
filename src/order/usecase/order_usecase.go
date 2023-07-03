package usecase

import (
	"context"
	"yandex-team.ru/bstask/model"
	"yandex-team.ru/bstask/order/handler/dto"
)

type OrderUsecase struct {
	orderRepo Repository
}

func NewOrderUsecase(repo Repository) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: repo,
	}
}

func (uc *OrderUsecase) GetOrder(ctx context.Context, orderId int64) (*dto.OrderDto, error) {
	entry, err := uc.orderRepo.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return orderModelToDto(entry), nil

}

func (uc *OrderUsecase) GetOrders(ctx context.Context, limit, offset int32) ([]*dto.OrderDto, error) {
	entry, err := uc.orderRepo.GetOrders(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return ordersSliceModelToDto(entry), nil
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, orders *dto.CreateOrderRequest) ([]*dto.OrderDto, error) {
	req := make([]*model.CreateOrder, 0, len(orders.Orders))
	for _, o := range orders.Orders {
		r := o.MapToModel()
		req = append(req, &r)
	}

	entry, err := uc.orderRepo.CreateOrders(ctx, req)
	if err != nil {
		return nil, err
	}

	return ordersSliceModelToDto(entry), nil
}

func (uc *OrderUsecase) CompleteOrders(ctx context.Context, orders *dto.CompleteOrderRequestDto) ([]*dto.OrderDto, error) {
	req := make([]*model.CompleteOrder, 0, len(orders.CompleteInfo))
	for _, o := range orders.CompleteInfo {
		r := model.CompleteOrder{
			CourierId:    o.CourierId,
			OrderId:      o.OrderId,
			CompleteTime: o.CompleteTime,
		}
		//r := o.MapToModel()
		req = append(req, &r)
	}
	entry, err := uc.orderRepo.UpdateOrders(ctx, req)
	if err != nil {
		return nil, err
	}

	return ordersSliceModelToDto(entry), nil
}

func ordersSliceModelToDto(entry []*model.Order) []*dto.OrderDto {
	resp := make([]*dto.OrderDto, 0, len(entry))
	for _, e := range entry {
		resp = append(resp, orderModelToDto(e))
	}
	return resp
}

func orderModelToDto(m *model.Order) *dto.OrderDto {
	if m == nil {
		return nil
	}
	return &dto.OrderDto{
		OrderId:       m.OrderId,
		Weight:        m.Weight,
		Regions:       m.Regions,
		DeliveryHours: m.DeliveryHours,
		Cost:          m.Cost,
		CompletedTime: m.CompletedTime,
	}
}
