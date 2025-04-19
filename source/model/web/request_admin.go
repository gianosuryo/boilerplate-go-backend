package web

type RequestAdmin struct {
	Name             string   `json:"name" form:"name" example:"example@bumame.com"`
	Email            string   `json:"email" form:"email" example:"example@bumame.com"`
	Password         string   `json:"password" form:"password" example:"asdf1234"`
	Role             string   `json:"role" form:"role" example:"b2b-sales"`
	Position         string   `json:"position" form:"position" example:"Manager"`
	CompanyClientIds []uint64 `json:"company_client_ids" form:"company_client_ids" example:"1,2"`
}

type QueryAdminB2B struct {
	Role string `query:"role"`
}
