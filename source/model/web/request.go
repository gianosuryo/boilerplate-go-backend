package web

type Request struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}

type QueryDashboard struct {
	DateStart string `query:"date_start" validate:"required" example:"2021-01-31"`
	DateEnd   string `query:"date_end" validate:"required" example:"2021-02-28"`
}
