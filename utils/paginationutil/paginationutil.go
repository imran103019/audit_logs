package paginationutil

// Page ...
type Page struct {
	Limit   uint `json:"per_page"`
	Current uint `json:"current_page"`
	Total   uint `json:"total"`
}

// Paginator ...
type Paginator interface {
	PageOffset() uint
	TotalPage() uint
	SetTotalPage(uint)
	PageLimit() uint
}

// PageOffset ...
func (p *Page) PageOffset() uint {
	if p.Current < 2 {
		return 0
	}
	return (p.Current - 1) * p.PageLimit()
}

// PageLimit ...
func (p *Page) PageLimit() uint {
	return p.Limit
}

// TotalPage ...
func (p *Page) TotalPage() uint {
	return p.Total
}

// SetTotalPage ...
func (p *Page) SetTotalPage(total uint) {
	p.Total = total
}

// NewPage ...
func NewPage() Page {
	return Page{
		Limit:   10,
		Current: 1,
		Total:   0,
	}
}

// NewPaginator ...
func NewPaginator(perPage, currentPage uint) Paginator {
	if perPage == 0 {
		perPage = 10
	}
	if currentPage == 0 {
		currentPage = 1
	}
	return &Page{Current: currentPage, Limit: perPage}
}