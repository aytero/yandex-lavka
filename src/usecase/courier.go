package usecase

import (
    "context"
    "fmt"
    "time"
    "yandex-team.ru/bstask/courier"
    "yandex-team.ru/bstask/model"
)

type CourierUsecase struct {
    courierRepo courier.Repository
}

func NewCourierUsecase(repo courier.Repository) *CourierUsecase {
    return &CourierUsecase{
        courierRepo: repo,
    }
}

func (uc *CourierUsecase) GetCourier(ctx context.Context, courierId int64) (*model.CourierDto, error) {
    entry, err := uc.courierRepo.GetById(ctx, courierId) // uc.BalanceRepo.GetById ; getUserByID
    if err != nil {
        //return nil, fmt.Errorf("CourierRepository - CreateCouriers: %w", err)
        return nil, err
    }
    return entry, nil
}

func (uc *CourierUsecase) GetCouriers(ctx context.Context, limit, offset int32) (*model.GetCouriersResponse, error) {
    entry, err := uc.courierRepo.GetCouriers(ctx, limit, offset)
    if err != nil {
        // todo logging
        return nil, err
    }
    resp := model.GetCouriersResponse{
        Couriers: entry,
        Limit:    limit,
        Offset:   offset,
    }
    return &resp, nil
}

func (uc *CourierUsecase) CreateCourier(ctx context.Context, couriers *model.CreateCourierRequest) ([]*model.CourierDto, error) {
    entry, err := uc.courierRepo.CreateCouriers(ctx, couriers)
    if err != nil {
        return nil, err
    }
    return entry, nil
}

func (uc *CourierUsecase) GetCourierMetaInfo(ctx context.Context, id int64, startDate, endDate time.Time) (*model.GetCourierMetaInfoResponse, error) {
    c, err := uc.courierRepo.GetById(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("CourierUsecase - GetMetaInfo: %w", err)
    }
    earnings, err := uc.courierRepo.GetEarnings(ctx, id, startDate, endDate)
    if err != nil {
        return nil, fmt.Errorf("CourierUsecase - GetMetaInfo: %w", err)
    }

    res := &model.GetCourierMetaInfoResponse{
        CourierId:    c.CourierId,
        CourierType:  c.CourierType,
        Regions:      c.Regions,
        WorkingHours: c.WorkingHours,
    }
    if len(earnings) == 0 {
        return res, nil
    }
    //coeff := getCoefficient(res.CourierType)
    res.Rating = calculateRating(startDate, endDate, ratingCoef[res.CourierType], int32(len(earnings)))
    res.Earnings = calculateEarnings(earnings, salaryCoef[res.CourierType])
    return res, nil
}

//type Base int

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

// todo enum
//func getCoefficient(courierType string) int32 {
//    switch courierType {
//    case "FOOT":
//        return 2
//    case "BIKE":
//        return 3
//    case "CAR":
//        return 4
//    default:
//        return 0
//    }
//}

/*
**Заработок рассчитывается по формуле:**
Заработок рассчитывается как сумма оплаты за каждый завершенный развоз в период с `start_date` (включая) до
`end_date` (исключая):
`sum = ∑(cost * C)`
`C`  — коэффициент, зависящий от типа курьера:
* пеший — 2
* велокурьер — 3
* авто — 4
Если курьер не завершил ни одного развоза, то рассчитывать и возвращать заработок не нужно.

**Рейтинг рассчитывается по формуле:**
Рейтинг рассчитывается следующим образом:
((число всех выполненных заказов с `start_date` по `end_date`) / (Количество часов между `start_date` и `end_date`)) * C
C - коэффициент, зависящий от типа курьера:
* пеший = 3
* велокурьер = 2
* авто - 1
Если курьер не завершил ни одного развоза, то рассчитывать и возвращать рейтинг не нужно.
*/

func calculateRating(startDate, endDate time.Time, coef int32, numOrders int32) int32 {
    hoursWorked := int32(endDate.Sub(startDate).Hours())
    return (numOrders / hoursWorked) * coef
}

func calculateEarnings(costs []int32, coef int32) int32 {
    var earnings int32
    for _, c := range costs {
        earnings += c * coef
    }
    return earnings
}
