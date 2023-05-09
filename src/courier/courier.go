package courier

import "context"

type Repository interface {
    GetById(ctx context.Context, userID int64) (*model.CourierDto, error)
    // other methods
}

type Usecase interface {
    GetCourier(ctx context.Context, userID int64) (*model.CourierDto, error)
    // other methods
}
