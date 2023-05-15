package courier

import (
    "context"
    "time"
    "yandex-team.ru/bstask/handler/dto"
    "yandex-team.ru/bstask/model"
)

type Repository interface {
    GetById(ctx context.Context, id int64) (*model.Courier, error)
    GetCouriers(ctx context.Context, limit, offset int32) ([]*model.Courier, error)
    CreateCouriers(ctx context.Context, couriers []*model.CreateCourier) ([]*model.Courier, error)
    GetEarnings(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, error)
}

type Usecase interface {
    GetCourier(ctx context.Context, userID int64) (*dto.CourierDto, error)
    GetCouriers(ctx context.Context, limit, offset int32) (*dto.GetCouriersResponse, error)
    CreateCourier(ctx context.Context, couriers *dto.CreateCourierRequest) (*dto.CreateCouriersResponse, error)
    GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*dto.GetCourierMetaInfoResponse, error)
}
