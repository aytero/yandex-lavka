package dto

import (
	"yandex-team.ru/bstask/model"
)

type CourierDto struct {
	CourierId    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type GetCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
	Limit    int32        `json:"limit"`
	Offset   int32        `json:"offset"`
}

type GetCourierMetaInfoResponse struct {
	CourierId    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
	Rating       int32    `json:"rating,omitempty"`
	Earnings     int32    `json:"earnings,omitempty"`
}

func ParseFromEntityMeta(entity *model.CourierMeta) *GetCourierMetaInfoResponse {
	return &GetCourierMetaInfoResponse{
		CourierId:    entity.CourierId,
		CourierType:  entity.CourierType,
		Regions:      entity.Regions,
		WorkingHours: entity.WorkingHours,
		Rating:       entity.Rating,
		Earnings:     entity.Earnings,
	}
}

func ParseFromEntitySlice(entity []*model.Courier) []CourierDto {
	resp := make([]CourierDto, 0, len(entity))
	for _, e := range entity {
		resp = append(resp, ParseFromEntity(e))
	}
	return resp
}

func ParseFromEntity(m *model.Courier) CourierDto {
	if m == nil {
		return CourierDto{}
	}
	return CourierDto{
		CourierId:    m.CourierId,
		CourierType:  m.CourierType,
		Regions:      m.Regions,
		WorkingHours: m.WorkingHours,
	}
}
