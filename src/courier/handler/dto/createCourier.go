package dto

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"yandex-team.ru/bstask/model"
)

const (
	foot string = "FOOT"
	bike        = "BIKE"
	car         = "CAR"
)

type CreateCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
}

type CreateCourierDto struct {
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers"`
}

func (c CreateCourierRequest) MapToModel() []*model.CreateCourier {
	reqC := make([]*model.CreateCourier, 0, len(c.Couriers))
	for _, cour := range c.Couriers {
		r := model.CreateCourier{CourierType: cour.CourierType, Regions: cour.Regions, WorkingHours: cour.WorkingHours}
		reqC = append(reqC, &r)
	}
	return reqC
}

func (c *CreateCourierRequest) Validate() error {
	for _, cour := range c.Couriers {
		if len(cour.WorkingHours) == 0 || len(cour.Regions) == 0 {
			return errors.New("json validation error")
		}
		if err := validateTime(cour.WorkingHours); err != nil {
			return fmt.Errorf("json validation error: %w", err)
		}
		if err := validateRegion(cour); err != nil {
			return fmt.Errorf("json validation error: %w", err)
		}
		if cour.CourierType != foot && cour.CourierType != bike && cour.CourierType != car {
			return errors.New("json validation error: invalid courier type")
		}
	}
	return nil
}

func validateRegion(cour CreateCourierDto) error {
	m := make(map[int32]struct{}, len(cour.Regions))
	for _, reg := range cour.Regions {
		if _, exists := m[reg]; exists {
			return errors.New("duplicate region")
		}
		m[reg] = struct{}{}
	}
	return nil
}

type timeInterval struct {
	start string
	end   string
}

func parseTimeInterval(workingHours string) (string, string) {
	interval := regexp.MustCompile(`-`).Split(workingHours, 2)
	return interval[0], interval[1]
}

func validateTime(workingHours []string) error {

	timeRegex := regexp.MustCompile(`^([0-1][0-9]|2[0-3]):[0-5][0-9]-([0-1][0-9]|2[0-3]):[0-5][0-9]$`)
	intervals := make([]timeInterval, len(workingHours))

	for _, interval := range workingHours {
		if err := timeRegex.MatchString(interval); !err {
			return errors.New("invalid time interval")
		}
		start, end := parseTimeInterval(interval)
		if start > end || start == end {
			return errors.New("invalid time interval")
		}
		intervals = append(intervals, timeInterval{start, end})
	}

	if len(intervals) > 1 {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].start < intervals[j].start
		})

		for i := 0; i < len(intervals)-1; i++ {
			if intervals[i].end > intervals[i+1].start {
				return errors.New("invalid time interval")
			}
		}
	}
	return nil
}
