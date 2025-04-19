package web

type RequestDailySlot struct {
	SlotQuota uint `json:"slot_quota" validate:"required" example:"10"`
}
