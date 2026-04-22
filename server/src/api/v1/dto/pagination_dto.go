package dto

type PaginationDto struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	PageCount int `json:"page_count"`
}
