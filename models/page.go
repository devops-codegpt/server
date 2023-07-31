package models

// PageInfo includes the Response result of query by page number
type PageInfo struct {
	PageNum      uint  `json:"pageNum" form:"pageNum"`
	PageSize     uint  `json:"pageSize" form:"pageSize"`
	Total        int64 `json:"total" form:"total"`
	NoPagination bool  `json:"noPagination" form:"noPagination"` // Do not use pagination
}

const (
	defaultPageSize = 10
	defaultPageNum  = 1
)

// GetLimit Calculate limit/offset for gorm query
func (p *PageInfo) GetLimit() (limit, offset int) {
	var pageSize, pageNum int64

	// The number of displayed items per page cannot be less than 1
	if p.PageSize < 1 {
		pageSize = defaultPageSize
	} else {
		pageSize = int64(p.PageSize)
	}
	// Page number cannot be less than 1
	if p.PageNum < 1 {
		pageNum = defaultPageNum
	} else {
		pageNum = int64(p.PageNum)
	}

	if p.Total > 0 && pageNum > p.Total {
		pageNum = p.Total
	}

	maxPageNum := p.Total/pageSize + 1
	if p.Total%pageSize == 0 {
		maxPageNum = p.Total / pageSize
	}
	if maxPageNum < 1 {
		maxPageNum = defaultPageNum
	}
	if pageNum > maxPageNum {
		pageNum = maxPageNum
	}

	limit = int(pageSize)
	offset = limit * int(pageNum-1)
	return limit, offset
}
