package web

type RequestB2BEquipmentConsumable struct {
	Name     string `json:"name" validate:"required" example:"Product Name"`
	Price    uint64 `json:"price" validate:"required" example:"10000000"`
	Quantity uint64 `json:"quantity" validate:"required" example:"100"`
	Unit     string `json:"unit" validate:"required" example:"pcs"`
	Type     string `json:"type" validate:"required" example:"equipment | consumable"`
}

type RequestB2BEquipmentConsumableChangeQuantity struct {
	Quantity uint64 `json:"quantity" validate:"required" example:"100"`
}
