package handler

import (
	"context"
	"time"
	"yandex-team.ru/bstask/model"
)

type Usecase interface {
	GetCourier(ctx context.Context, userID int64) (*model.Courier, error)
	GetCouriers(ctx context.Context, limit, offset int32) ([]*model.Courier, error)
	CreateCourier(ctx context.Context, couriers []*model.CreateCourier) ([]*model.Courier, error)
	GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*model.CourierMeta, error)
}
