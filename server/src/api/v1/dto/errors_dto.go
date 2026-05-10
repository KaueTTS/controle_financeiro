package dto

type ErrorDto struct {
	Message     string          `json:"message"`
	CodeMessage string          `json:"codeMessage,omitempty"`
	Details     []FieldErrorDto `json:"details,omitempty"`
}

type FieldErrorDto struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}
