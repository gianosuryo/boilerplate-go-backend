package web

type RequestBumameScheduleRandomizer struct {
	Month int `json:"month" form:"month" validate:"required"`
	Year  int `json:"year" form:"year" validate:"required"`

	MaxHQPerDay               int `json:"max_hq_per_day" form:"max_hq_per_day"`
	MaxDayOffPerNakesPerMonth int `json:"max_day_off_per_nakes_per_month" form:"max_day_off_per_nakes_per_month"`
	MaxOffNakesCountPerDay    int `json:"max_off_nakes_count_per_day" form:"max_off_nakes_count_per_day"`
}

type RequestBumameScheduleSubmit struct {
	Month            uint   `json:"month" form:"month" validate:"required"`
	Year             uint   `json:"year" form:"year" validate:"required"`
	CalendarDataJson string `json:"calendar_data_json" form:"calendar_data_json" validate:"required"`
}
