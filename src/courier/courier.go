package courier

import (
	"context"
	"yandex-team.ru/bstask/model"
)

type Repository interface {
	GetById(ctx context.Context, userID int64) (*model.CourierDto, error)
	// other methods
}

type Usecase interface {
	GetCourier(ctx context.Context, userID int64) (*model.CourierDto, error)
	// other methods
}
