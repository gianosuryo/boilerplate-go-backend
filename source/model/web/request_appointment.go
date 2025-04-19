package web

type RequestBumameAppointment struct {
	CompanyClientId       uint64  `json:"company_client_id" form:"company_client_id" validate:"required" example:"1"`
	SalesAdminId          uint64  `json:"sales_admin_id" form:"sales_admin_id" validate:"required" example:"1"`
	ProjectType           string  `json:"project_type" form:"project_type" validate:"required" example:"On-site"`
	ExecutionType         string  `json:"execution_type" form:"execution_type" validate:"required" example:"In-house"`
	NumberOfPatient       uint64  `json:"number_of_patient" form:"number_of_patient" validate:"required" example:"1"`
	AppointmentNote       string  `json:"appointment_note" form:"appointment_note" example:"Appointment note"`
	TotalPricePreDiscount uint64  `json:"total_price_pre_discount" form:"total_price_pre_discount" validate:"required" example:"10000000"`
	TotalDiscount         uint64  `json:"total_discount" form:"total_discount" validate:"required" example:"500000"`
	DiscountPercentage    float32 `json:"discount_percentage" form:"discount_percentage" validate:"required" example:"5"`
	TotalAmount           uint64  `json:"total_amount" form:"total_amount" validate:"required" example:"9500000"`
}

type RequestBumameAppointmentHeldDate struct {
	HeldDateStringArr []string `json:"held_date_string_arr" form:"held_date_string_arr" validate:"required" example:"2025-01-01"`
}

type RequestBumameAppointmentDetailProduct struct {
	B2BProductId uint64 `json:"b2b_product_id" form:"b2b_product_id" validate:"required" example:"1"`
	Quantity     uint64 `json:"quantity" form:"quantity" validate:"required" example:"1"`
	Price        uint64 `json:"price" form:"price" validate:"required" example:"10000000"`
	Discount     uint64 `json:"discount" form:"discount" validate:"required" example:"500000"`
	Total        uint64 `json:"total" form:"total" validate:"required" example:"9500000"`
}

type RequestBumameAppointmentDetailPackage struct {
	PackageId uint64 `json:"package_id" form:"package_id" validate:"required" example:"1"`
	Quantity  uint64 `json:"quantity" form:"quantity" validate:"required" example:"1"`
	Price     uint64 `json:"price" form:"price" validate:"required" example:"10000000"`
	Discount  uint64 `json:"discount" form:"discount" validate:"required" example:"500000"`
	Total     uint64 `json:"total" form:"total" validate:"required" example:"9500000"`
}

type RequestBumameAppointmentDetailEquipmentConsumable struct {
	B2BEquipmentConsumableId uint64 `json:"b2b_equipment_consumable_id" form:"b2b_equipment_consumable_id" validate:"required" example:"1"`
	Quantity                 uint64 `json:"quantity" form:"quantity" validate:"required" example:"1"`
}

type RequestBumameAppointmentChangeStatus struct {
	AppointmentStatus string `json:"appointment_status" form:"appointment_status" validate:"required" enums:"waiting_for_approval,approved,in_progress,post_progress,completed,cancelled,rejected,closed" example:"waiting_for_approval"`

	IsApprovedHeadOfSales uint `json:"is_approved_head_of_sales" form:"is_approved_head_of_sales" example:"1"`
	IsApprovedHeadOfOps   uint `json:"is_approved_head_of_ops" form:"is_approved_head_of_ops" example:"1"`
	IsApprovedCeo         uint `json:"is_approved_ceo" form:"is_approved_ceo" example:"1"`
}

type RequestBumameAppointmentManpowerReady struct {
	IsManpowerReady uint `json:"is_manpower_ready" form:"is_manpower_ready" example:"1"`
}
