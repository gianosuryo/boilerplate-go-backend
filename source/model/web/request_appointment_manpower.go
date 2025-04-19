package web

type RequestAppointmentManpower struct {
	AdminId uint64 `json:"admin_id" validate:"required" example:"1"`
	Role    string `json:"role" validate:"required" example:"admin" enums:"admin,ttv,visus,dokter,plebo"`
}
