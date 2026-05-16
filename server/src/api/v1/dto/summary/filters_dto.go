package dto_summary

type SummaryFilterDto struct {
	Search   string `query:"search"`
	Type     string `query:"type"`
	Category string `query:"category"`
}
