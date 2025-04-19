package web

type RequestCompanyClient struct {
	Name         string `json:"name" validate:"required" example:"PT. Astra International"`
	IndustryType string `json:"industry_type" validate:"required" example:"Automotive"`
	Address      string `json:"address" validate:"required" example:"Jl. Jendral Sudirman Kav. 52-53, RT.5/RW.3, Senayan, Kebayoran Baru, Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12190"`
	Phone        string `json:"phone" validate:"required" example:"021-567890"`
	Email        string `json:"email" validate:"required,email" example:"info@astra.com"`

	ContactPersonPhone    string `json:"contact_person_phone" example:"021-567890"`
	ContactPersonEmail    string `json:"contact_person_email" validate:"email" example:"contact@astra.com"`
	ContactPersonPosition string `json:"contact_person_position" example:"Manager"`
	ContactPersonName     string `json:"contact_person_name" example:"John Doe"`
}
