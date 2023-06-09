package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	mockCourier "yandex-team.ru/bstask/courier/repository/mocks"
	"yandex-team.ru/bstask/model"
)

func TestGetCouriers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mockCourier.NewMockRepository(ctl)

	ctx := context.Background()

	mockRepo := []*model.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	expected := []*model.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}
	/*
		expected := dto.GetCouriersResponse{
			Couriers: []dto.CourierDto{
				{
					CourierId:    1,
					CourierType:  "FOOT",
					Regions:      []int32{2, 4},
					WorkingHours: []string{"12:34:00", "12:12:00"},
				},
			},
			Limit:  int32(1),
			Offset: int32(0),
		}
	*/

	// Validity check
	repo.EXPECT().GetCouriers(ctx, int32(1), int32(0)).Return(mockRepo, nil).Times(1)

	service := NewCourierUsecase(repo)
	orders, err := service.GetCouriers(ctx, int32(1), int32(0))
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	// Invalidity check
	errDb := errors.New("db is down")
	repo.EXPECT().GetCouriers(ctx, int32(1), int32(0)).Return(nil, errDb).Times(1)

	_, err = service.GetCouriers(ctx, int32(1), int32(0))
	require.Error(t, err)

}

func TestCreateCourier(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mockCourier.NewMockRepository(ctl)

	ctx := context.Background()

	mockRepo := []*model.CreateCourier{
		{
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	expected := []*model.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	retur := []*model.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	//проверка на валидность
	repo.EXPECT().CreateCouriers(ctx, mockRepo).Return(retur, nil).Times(1)

	service := NewCourierUsecase(repo)
	orders, err := service.CreateCourier(ctx, mockRepo)
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().CreateCouriers(ctx, mockRepo).Return(nil, errDb).Times(1)

	_, err = service.CreateCourier(ctx, mockRepo)
	require.Error(t, err)

}

func TestGetCourierById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mockCourier.NewMockRepository(ctl)

	ctx := context.Background()

	retur := &model.Courier{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}

	/*
		expected := &dto.CourierDto{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		}
	*/
	expected := &model.Courier{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}

	repo.EXPECT().GetById(ctx, int64(1)).Return(retur, nil).Times(1)

	service := NewCourierUsecase(repo)
	orders, err := service.GetCourier(ctx, int64(1))
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().GetById(ctx, int64(1)).Return(nil, errDb).Times(1)

	_, err = service.GetCourier(ctx, int64(1))
	require.Error(t, err)
}

func TestGetCourierMetaInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mockCourier.NewMockRepository(ctl)

	ctx := context.Background()

	retur := []model.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "BIKE",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "CAR",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}
	var r1 int32 = 0
	var e1 int32 = 3800

	var r2 int32 = 0
	var e2 int32 = 5700

	var r3 int32 = 0
	var e3 int32 = 7600

	var r4 int32 = 0
	var e4 int32 = 0

	//expected := []dto.GetCourierMetaInfoResponse{
	expected := []model.CourierMeta{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       r1,
			Earnings:     e1,
		},
		{
			CourierId:    1,
			CourierType:  "BIKE",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       r2,
			Earnings:     e2,
		},
		{
			CourierId:    1,
			CourierType:  "CAR",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       r3,
			Earnings:     e3,
		},
		{
			CourierId:    1,
			CourierType:  "",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       r4,
			Earnings:     e4,
		},
	}

	cases := []string{
		"test_1", "test_2", "test_3", "test_4",
	}
	t1, _ := time.Parse("2006-01-02", "2023-01-01")
	t2, _ := time.Parse("2006-01-02", "2023-01-02")

	cost := []int32{500, 1000, 400}

	for i, name := range cases {
		t.Run(name, func(t *testing.T) {
			repo.EXPECT().GetById(ctx, int64(1)).Return(&retur[i], nil).Times(1)
			repo.EXPECT().GetEarnings(ctx, int64(1), t1, t2).Return(cost, nil).Times(1)
			//repo.EXPECT().GetEarnings(ctx, int64(1), t1, t2).Return(cost, &retur[i], nil).Times(1)

			service := NewCourierUsecase(repo)
			orders, _ := service.GetCourierMetaInfo(ctx, int64(1), t1, t2)
			require.Equal(t, &expected[i], orders)
			//require.Equal(t, &expected[i], orders)
		})
	}

	repo.EXPECT().GetById(ctx, int64(1)).Return(&retur[0], nil).Times(1)
	repo.EXPECT().GetEarnings(ctx, int64(1), time.Time{}, time.Time{}).Return(nil, nil).Times(1)
	//repo.EXPECT().GetEarnings(ctx, int64(1), time.Time{}, time.Time{}).Return(nil, &retur[0], nil).Times(1)

	service := NewCourierUsecase(repo)
	orders, _ := service.GetCourierMetaInfo(ctx, int64(1), time.Time{}, time.Time{})
	//newEx := &dto.GetCourierMetaInfoResponse{
	newEx := &model.CourierMeta{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}
	require.Equal(t, newEx, orders)
}

func TestCourierUsecase_GetCourierMetaInfo(t *testing.T) {
	type fields struct {
		courierRepo Repository
	}
	type args struct {
		ctx       context.Context
		id        int64
		startDate time.Time
		endDate   time.Time
	}

	ret := &model.Courier{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}
	expected := &model.CourierMeta{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
		Rating:       0,
		Earnings:     3800,
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockRepo := mockCourier.NewMockRepository(ctl)
	ctx := context.Background()
	t1, _ := time.Parse("2006-01-02", "2023-01-01")
	t2, _ := time.Parse("2006-01-02", "2023-01-02")
	cost := []int32{500, 1000, 400}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CourierMeta
		wantErr bool
	}{
		{"1", fields{mockRepo}, args{ctx, 1, t1, t2}, expected, false},
		//{name: "1", fields: fields{mockRepo}, args: args{ctx, 1,"12:34:00", "12:12:00" }, want: expected, wantErr: true},
		//{"3 should be Foo", args{3}, "Foo"},
		//{"1 is not Foo", args{1}, "1"},
		//{"0 should be Foo", args{0}, "Foo"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CourierUsecase{
				courierRepo: tt.fields.courierRepo,
			}
			mockRepo.EXPECT().GetById(ctx, int64(1)).Return(ret, nil).Times(1)
			mockRepo.EXPECT().GetEarnings(ctx, int64(1), t1, t2).Return(cost, nil).Times(1)
			got, err := uc.GetCourierMetaInfo(tt.args.ctx, tt.args.id, tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCourierMetaInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCourierMetaInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
