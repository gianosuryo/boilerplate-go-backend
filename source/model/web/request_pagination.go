package web

type Filter map[string]string

type RequestPaginationNumber struct {
	Page  int    `json:"page" query:"page" default:"1"`
	Limit int    `json:"limit" query:"limit" default:"1000"`
	Sort  string `json:"sort" query:"sort" default:"id desc"`
	Search string `json:"search" query:"search"`

	CategoryId string `json:"health_service_category_id" query:"health_service_category_id"`
	Keyword    string `json:"keyword" query:"keyword"`
	// Filter Filter `json:"filter" query:"filter"`
}

type Person struct {
	Name     string   `query:"name"`
	Pass     string   `query:"pass"`
	Products []string `query:"products"`
}

type RequestPaginationInfinite struct {
	Page   string `json:"page" query:"page"`
	Limit  int    `json:"limit" query:"limit" `
	Sort   string `json:"sort" query:"sort"`
	Filter Filter `json:"filter"`
}
