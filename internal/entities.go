package internal

type RequiredField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidateReturn struct {
	Message string          `json:"message"`
	Errors  []RequiredField `json:"errors"`
}
