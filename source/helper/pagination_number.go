package helper

import (
	"math"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type PaginationNumber struct {
	Type           string            `json:"type"`
	Page           int               `json:"page"`
	PerPage        int               `json:"per_page"`
	TotalRows      int64             `json:"total_rows"`
	TotalPages     int               `json:"total_pages"`
	NumberingStart int64             `json:"numbering_start"`
	Sort           string            `json:"sort"`
	Search         string            `json:"search"`
	MainTableName  string            `json:"table_name,omitempty"`
	Query          map[string]string `json:"query,omitempty"`
}

type MongoPagination struct {
	limit int64
	page  int64
}

func (p *PaginationNumber) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *PaginationNumber) GetLimit() int {
	if p.PerPage == 0 {
		p.PerPage = 1000
	}
	return p.PerPage
}
func (p *PaginationNumber) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *PaginationNumber) GetSort() string {
	if p.Sort == "" {
		if p.MainTableName != "" {
			p.Sort = p.MainTableName + "."

		}
		p.Sort = p.Sort + "id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *PaginationNumber, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pagination.Type = "number"
	if pagination.TotalRows == 0 {
		db.Model(value).Count(&pagination.TotalRows)
	}

	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	if pagination.GetPage() == 1 {
		pagination.NumberingStart = int64(pagination.GetPage())
	} else {
		pagination.NumberingStart = int64(pagination.Page*pagination.PerPage - (pagination.PerPage - 1))
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func MongoPaginate(pagination *PaginationNumber) *MongoPagination {
	pagination.GetLimit()

	if pagination.GetPage() == 1 {
		pagination.NumberingStart = int64(pagination.GetPage())
	} else {
		pagination.NumberingStart = int64((pagination.Page * pagination.PerPage) - (pagination.PerPage - 1))
	}

	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return &MongoPagination{
		limit: int64(pagination.PerPage),
		page:  int64(pagination.Page),
	}
}

func (mp *MongoPagination) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := (mp.page * mp.limit) - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}
