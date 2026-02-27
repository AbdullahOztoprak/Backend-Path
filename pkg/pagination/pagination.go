package pagination

import (
	"math"
	"net/http"
	"strconv"
)

// Pagination holds the pagination information
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// NewPagination creates a new Pagination instance
func NewPagination(page, pageSize, totalItems int) *Pagination {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

// GetPaginationFromRequest extracts pagination parameters from the request
func GetPaginationFromRequest(r *http.Request) (int, int) {
	page := 1
	pageSize := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	if ps := r.URL.Query().Get("page_size"); ps != "" {
		if parsedPageSize, err := strconv.Atoi(ps); err == nil {
			pageSize = parsedPageSize
		}
	}

	return page, pageSize
}