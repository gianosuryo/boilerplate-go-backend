package web

type RequestHealthServiceCategory struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhotoUrl    string `json:"photo_url" form:"photo_url"`
	Description string `json:"description" form:"description"`
}
