package handler

import "github.com/fanesz/bindator/internal"

func Response(ok bool, message string) *internal.ValidateReturn {
	return &internal.ValidateReturn{
		Ok:      ok,
		Message: message,
	}
}
