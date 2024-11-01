package internal

type RequiredField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidateReturn struct {
	Ok      bool            `json:"ok,omitempty"`
	Message string          `json:"message"`
	Errors  []RequiredField `json:"errors,omitempty"`
}
