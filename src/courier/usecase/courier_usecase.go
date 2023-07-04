package usecase

import (
	"context"
	"fmt"
	"time"
	"yandex-team.ru/bstask/model"
)

type CourierUsecase struct {
	courierRepo Repository
}

func NewCourierUsecase(repo Repository) *CourierUsecase {
	return &CourierUsecase{
		courierRepo: repo,
	}
}

func (uc *CourierUsecase) GetCourier(ctx context.Context, courierId int64) (*model.Courier, error) {
	entry, err := uc.courierRepo.GetById(ctx, courierId)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, nil
	}
	return entry, nil
}

func (uc *CourierUsecase) GetCouriers(ctx context.Context, limit, offset int32) ([]*model.Courier, error) {
	entry, err := uc.courierRepo.GetCouriers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *CourierUsecase) CreateCourier(ctx context.Context, couriers []*model.CreateCourier) ([]*model.Courier, error) {

	entry, err := uc.courierRepo.CreateCouriers(ctx, couriers)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *CourierUsecase) GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*model.CourierMeta, error) {
	c, err := uc.courierRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("CourierUsecase - GetMetaInfo: %w", err)
	}
	earnings, err := uc.courierRepo.GetEarnings(ctx, id, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("CourierUsecase - GetMetaInfo: %w", err)
	}

	if c == nil {
		return &model.CourierMeta{}, nil
	}
	res := &model.CourierMeta{
		CourierId:    c.CourierId,
		CourierType:  c.CourierType,
		Regions:      c.Regions,
		WorkingHours: c.WorkingHours,
	}
	if len(earnings) == 0 {
		return res, nil
	}
	res.Rating = uc.calculateRating(startDate, endDate, ratingCoef[res.CourierType], int32(len(earnings)))
	res.Earnings = uc.calculateEarnings(earnings, salaryCoef[res.CourierType])
	return res, nil
}

var ratingCoef = map[string]int32{
	"FOOT": 3,
	"BIKE": 2,
	"CAR":  1,
}

var salaryCoef = map[string]int32{
	"FOOT": 2,
	"BIKE": 3,
	"CAR":  4,
}

func (uc CourierUsecase) calculateRating(startDate, endDate time.Time, coef int32, numOrders int32) int32 {
	hoursWorked := endDate.Sub(startDate).Hours()
	return (numOrders / int32(hoursWorked)) * coef
}

func (uc CourierUsecase) calculateEarnings(costs []int32, coef int32) int32 {
	var earnings int32
	for _, c := range costs {
		earnings += c * coef
	}
	return earnings
}
