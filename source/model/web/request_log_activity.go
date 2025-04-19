package web

type RequestLogActivity struct {
	LogType string `json:"log_type" validate:"required"`
	LogData string `json:"log_data" validate:"required"`
	AdminId uint64 `json:"admin_id" validate:"required"`
}
