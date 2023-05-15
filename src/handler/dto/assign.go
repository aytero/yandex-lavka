package dto

// todo where does these go?

type OrderAssignResponse struct {
    Date     string                `json:"date"`
    Couriers []CouriersGroupOrders `json:"couriers"`
}

type CouriersGroupOrders struct {
    CourierId int64         `json:"courier_id"`
    Orders    []GroupOrders `json:"orders"`
}
