package web

import "time"

type RequestAskAvailableHealthService struct {
	AppointmentDateTime string `json:"appointment_date_time" form:"appointment_date_time"` // format: YYYY-MM-DD HH:MM:SS
}

type RequestGetUserAppointmentByPhoneNumber struct {
	PhoneNumber string `query:"phone_number" validate:"required"`
}

type RequestCreateSimpleUserAppointment struct {
	Name            string `json:"name" form:"name" validate:"required"`
	Address         string `json:"address" form:"address" validate:"required"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" validate:"required"`
	HealthServiceId uint64 `json:"health_service_id" form:"health_service_id" validate:"required"`
	AppointmentDate string `json:"appointment_date" form:"appointment_date" validate:"required"` // format: YYYY-MM-DD
	AppointmentTime string `json:"appointment_time" form:"appointment_time" validate:"required"` // format: HH:MM:SS

	UserId                       uint64    `json:"user_id"`
	FormattedAppointmentDateTime time.Time `json:"formatted_appointment_date_time"`
}
