package handler

import (
	"context"
	"time"
	"yandex-team.ru/bstask/courier/handler/dto"
)

type Usecase interface {
	GetCourier(ctx context.Context, userID int64) (*dto.CourierDto, error)
	GetCouriers(ctx context.Context, limit, offset int32) (*dto.GetCouriersResponse, error)
	CreateCourier(ctx context.Context, couriers *dto.CreateCourierRequest) (*dto.CreateCouriersResponse, error)
	GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*dto.GetCourierMetaInfoResponse, error)
}
