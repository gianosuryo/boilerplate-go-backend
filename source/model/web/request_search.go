package web

type RequestSearch struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"required"`
}
