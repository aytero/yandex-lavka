package usecase

import (
	"context"
	"time"
	"yandex-team.ru/bstask/model"
)

type Repository interface {
	GetById(ctx context.Context, id int64) (*model.Courier, error)
	GetCouriers(ctx context.Context, limit, offset int32) ([]*model.Courier, error)
	CreateCouriers(ctx context.Context, couriers []*model.CreateCourier) ([]*model.Courier, error)
	GetEarnings(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, error)
}
