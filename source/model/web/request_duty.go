package web

type RequestDuty struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Code      string `json:"code" form:"code" validate:"required"`
	ColorCode string `json:"color_code" form:"color_code" validate:"required"`
}
