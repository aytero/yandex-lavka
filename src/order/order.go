package order

import "context"

type Repository interface {
    GetById(ctx context.Context, userID int64) (*model.OrderDto, error)
    // other methods
}

type Usecase interface {
    GetOrder(ctx context.Context, userID int64) (*model.OrderDto, error)
    // other methods
}
