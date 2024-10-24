package bindator

import (
	"io"
	"reflect"

	"github.com/fanesz/bindator/internal"
	"github.com/gin-gonic/gin"
)

func BindBody(ctx *gin.Context, req interface{}) (ok bool, res *internal.ValidateReturn) {
	if err := ctx.ShouldBindJSON(req); err != nil {
		if err == io.EOF {
			return false, &internal.ValidateReturn{
				Message: "Invalid JSON data: unexpected end of JSON input",
			}
		}
		return false, &internal.ValidateReturn{
			Message: err.Error(),
		}
	}

	res, err := internal.ValidateRequest(&req, "body")
	if err != nil {
		return false, &internal.ValidateReturn{
			Message: err.Error(),
		}
	}
	if res != nil {
		return false, res
	}

	return true, nil
}

func BindBodies(ctx *gin.Context, obj interface{}) (ok bool, res *internal.ValidateReturn) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := ctx.ShouldBindJSON(field); err != nil {
			if err == io.EOF {
				return false, &internal.ValidateReturn{
					Message: "Invalid JSON data: unexpected end of JSON input",
				}
			}
			return false, &internal.ValidateReturn{
				Message: err.Error(),
			}
		}
		res, err := internal.ValidateRequest(&field, "body")
		if err != nil {
			return false, &internal.ValidateReturn{
				Message: err.Error(),
			}
		}
		if res != nil {
			return false, res
		}
	}
	return true, nil
}
