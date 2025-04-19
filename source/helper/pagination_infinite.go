package helper

type PaginationInfinite struct {
	NextPage      string            `json:"next_page"` // For infinite pagination
	Type          string            `json:"type"`
	PerPage       int               `json:"per_page"`
	MainTableName string            `json:"table_name,omitempty"`
	Query         map[string]string `json:"query,omitempty"`
}

func (p *PaginationInfinite) SetPerPage(perPage int) {
	p.PerPage = perPage
}
func (p *PaginationInfinite) GetPerPage() int {
	p.Type = "infinite"
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	return p.PerPage
}
func (p *PaginationInfinite) SetNextPage(nextPage string) {
	p.NextPage = nextPage
}
