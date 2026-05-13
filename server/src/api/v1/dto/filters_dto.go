package dto

type FilterDto struct {
	Search   string `query:"search"`
	Type     string `query:"type"`
	Category string `query:"category"`
	Page     int    `query:"page"`
	PerPage  int    `query:"perPage"`
}
