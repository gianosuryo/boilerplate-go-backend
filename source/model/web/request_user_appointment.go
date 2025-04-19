package web

import "time"

type RequestBumameUserAppointment struct {
	Name                string  `json:"name" form:"name" validate:"required"`
	Email               string  `json:"email" form:"email" validate:"required"`
	Phone               string  `json:"phone" form:"phone" validate:"required"`
	Address             string  `json:"address" form:"address" validate:"required"`
	Longitude           float64 `json:"longitude" form:"longitude" validate:"required"`
	Latitude            float64 `json:"latitude" form:"latitude" validate:"required"`
	AppointmentDateTime string  `json:"appointment_date_time" form:"appointment_date_time" validate:"required"`
	AppointmentNote     string  `json:"appointment_note" form:"appointment_note"`
	HealthServiceIds    []uint  `json:"health_service_ids" form:"health_service_ids" validate:"required"`
	MedicalRecordLog    string  `json:"medical_record_log" form:"medical_record_log"`

	PrescreeningResultJson string `json:"prescreening_result_json" form:"prescreening_result_json"` // array of json (question, answer)

	// no effect in request
	FormattedAppointmentDateTime time.Time `json:"formatted_appointment_date_time" form:"formatted_appointment_date_time"`
	UserId                       uint64    `json:"user_id" form:"user_id"`
}

type RequestBumameUserAppointmentV2 struct {
	Name                string                   `json:"name" form:"name" validate:"required"`
	Email               string                   `json:"email" form:"email" validate:"required"`
	Phone               string                   `json:"phone" form:"phone" validate:"required"`
	Address             string                   `json:"address" form:"address" validate:"required"`
	Gender              string                   `json:"gender" form:"gender" validate:"required"`
	BirthDate           string                   `json:"birth_date" form:"birth_date" validate:"required"`
	Nik                 string                   `json:"nik" form:"nik" validate:"required"`
	KnowingBumame       string                   `json:"knowing_bumame" form:"knowing_bumame" validate:"required"`
	Longitude           float64                  `json:"longitude" form:"longitude" validate:"required"`
	Latitude            float64                  `json:"latitude" form:"latitude" validate:"required"`
	AppointmentDateTime string                   `json:"appointment_date_time" form:"appointment_date_time" validate:"required"`
	AppointmentNote     string                   `json:"appointment_note" form:"appointment_note"`
	SelectedAdminId     uint64                   `json:"selected_admin_id" form:"selected_admin_id" validate:"required"`
	PatientDataJson     []RequestPatientDataJson `json:"patient_data_json" form:"patient_data_json" validate:"required"`

	MedicalRecordLog string `json:"medical_record_log" form:"medical_record_log"`

	// no effect in request
	FormattedAppointmentDateTime time.Time `json:"formatted_appointment_date_time" form:"formatted_appointment_date_time"`
	UserId                       uint64    `json:"user_id" form:"user_id"`
}

type RequestPatientDataJson struct {
	Name                   string                          `json:"name" form:"name" validate:"required"`
	Email                  string                          `json:"email" form:"email"`
	Phone                  string                          `json:"phone" form:"phone"`
	HealthServiceDataJson  []RequestHealthServiceDataJson  `json:"health_service_data_json" form:"health_service_data_json"`
	PrescreeningResultJson []RequestPrescreeningResultJson `json:"prescreening_result_json" form:"prescreening_result_json"`
}

type RequestHealthServiceDataJson struct {
	HealthServiceId uint64  `json:"health_service_id" form:"health_service_id" validate:"required"`
	Quantity        uint64  `json:"quantity" form:"quantity" validate:"required"`
	Price           uint64  `json:"price" form:"price"`
	HourDuration    float32 `json:"hour_duration" form:"hour_duration"`

	// no effect in request
	HealthServiceName string `json:"health_service_name" form:"health_service_name"`
	Equipment         string `json:"equipment" form:"equipment"`
	Consumables       string `json:"consumables" form:"consumables"`
}

type RequestPrescreeningResultJson struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer   string `json:"answer" form:"answer" validate:"required"`
}

type RequestBumameUserAppointmentSetStatus struct {
	AppointmentStatus string  `json:"appointment_status" form:"appointment_status" validate:"required"`
	AdminId           uint64  `json:"admin_id" form:"admin_id" default:"0"`
	Latitude          float64 `json:"latitude" form:"latitude"`
	Longitude         float64 `json:"longitude" form:"longitude"`
}

type RequestBumameUserAppointmentUpdatePayment struct {
	BankName          string `json:"bank_name" form:"bank_name" validate:"required"`
	BankAccountName   string `json:"bank_account_name" form:"bank_account_name" validate:"required"`
	BankAccountNumber string `json:"bank_account_number" form:"bank_account_number" validate:"required"`
	DoctorTeleConsultation string `json:"doctor_tele_consultation" form:"doctor_tele_consultation"`
	InternalNote string `json:"internal_note" form:"internal_note"`
	ReceiptFileUrl    string `json:"receipt_file_url" form:"receipt_file_url"`
}

type PrescreeningResultJsonInterface struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type RequestChangeNakes struct {
	AdminId uint64 `json:"admin_id" validate:"required"`
}

type QueryGetUserAppointment struct {
	Status           string `query:"status"`
	OrderId          string `query:"order_id"`
	HealthServiceId  uint64 `query:"service_id"`
	HealthWorkerName string `query:"health_worker_name"`
	OrderDateStart   string `query:"order_date_range"`
	OrderDateEnd     string `query:"order_date_range"`
	AppointmentDate  string `query:"appointment_date"`
}

type RequestCheckAvailableHourlySchedule struct {
	HealthServiceIds []uint64 `json:"health_service_ids" validate:"required"`
	AppointmentDate  string   `json:"appointment_date_time" validate:"required"`
	Longitude        float64  `json:"longitude" validate:"required"`
	Latitude         float64  `json:"latitude" validate:"required"`
}
