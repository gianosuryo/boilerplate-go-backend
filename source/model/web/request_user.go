package web

type RequestUser struct {
	Name      string  `json:"name" form:"name"`
	Phone     string  `json:"phone" form:"phone"`
	Address   string  `json:"address" form:"address"`
	Email     string  `json:"email" form:"email"`
	Longitude float64 `json:"longitude" form:"longitude"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Nik       string  `json:"nik" form:"nik"`
	BirthDate string  `json:"birth_date" form:"birth_date"`
	Gender    string  `json:"gender" form:"gender"`
	KnowingBumame string  `json:"knowing_bumame" form:"knowing_bumame"`
}
