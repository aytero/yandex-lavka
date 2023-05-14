package courier

import (
    "context"
    "time"
    "yandex-team.ru/bstask/model"
)

type Repository interface {
    GetById(ctx context.Context, id int64) (*model.CourierDto, error)
    GetCouriers(ctx context.Context, limit, offset int32) ([]model.CourierDto, error)
    CreateCouriers(ctx context.Context, ords *model.CreateCourierRequest) ([]*model.CourierDto, error)
    GetEarnings(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, error)
}

type Usecase interface {
    GetCourier(ctx context.Context, userID int64) (*model.CourierDto, error)
    GetCouriers(ctx context.Context, limit, offset int32) (*model.GetCouriersResponse, error)
    CreateCourier(ctx context.Context, couriers *model.CreateCourierRequest) ([]*model.CourierDto, error)
    GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*model.GetCourierMetaInfoResponse, error)
}
