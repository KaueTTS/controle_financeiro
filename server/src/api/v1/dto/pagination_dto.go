package dto

type PaginationDto struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"perPage"`
	PageCount int   `json:"pageCount"`
	Total     int64 `json:"total"`
}
