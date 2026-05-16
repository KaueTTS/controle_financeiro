package dto_shared

type ErrorDto struct {
	Message     string           `json:"message"`
	CodeMessage string           `json:"codeMessage,omitempty"`
	Details     []DetailErrorDto `json:"details,omitempty"`
}

type DetailErrorDto struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}
