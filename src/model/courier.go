package model

type Courier struct {
	CourierId    int64    `json:"courier_id" db:"courier_id"`
	CourierType  string   `json:"courier_type" db:"courier_type"`
	Regions      []int32  `json:"regions" db:"regions"`
	WorkingHours []string `json:"working_hours" db:"working_hours"`
	//Rating       int32    `db:"rating"`
	//Earnings     int32    `db:"earnings"`
}

type CourierMeta struct {
	CourierId    int64    `json:"courier_id" db:"courier_id"`
	CourierType  string   `json:"courier_type" db:"courier_type"`
	Regions      []int32  `json:"regions" db:"regions"`
	WorkingHours []string `json:"working_hours" db:"working_hours"`
	Rating       int32    `db:"rating"`
	Earnings     int32    `db:"earnings"`
}

type CreateCourier struct {
	CourierType  string   `json:"courier_type" db:"courier_type"`
	Regions      []int32  `json:"regions" db:"regions"`
	WorkingHours []string `json:"working_hours" db:"working_hours"`
}
