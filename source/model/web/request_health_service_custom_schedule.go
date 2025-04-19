package web

type RequestHealthServiceCustomSchedule struct {
	CustomScheduleType string   `json:"custom_schedule_type"; form:"custom_schedule_type"` // weekly, specific_date
	DayOfWeek          int      `json:"day_of_week"; form:"day_of_week"`                   // 1-7
	SpecificDate       string   `json:"specific_date"; form:"specific_date"`
	ServiceHours       []string `json:"service_hours"; form:"service_hours"`
}
