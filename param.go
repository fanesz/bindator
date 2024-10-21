package bindator

import (
	"reflect"

	"github.com/fanesz/bindator/internal"
	"github.com/gin-gonic/gin"
)

func BindParam(ctx *gin.Context, req interface{}) (ok bool, res *internal.ValidateReturn) {
	if err := ctx.ShouldBindQuery(req); err != nil {
		return false, &internal.ValidateReturn{
			Message: err.Error(),
		}
	}

	res, err := internal.ValidateRequest(&req, "param")
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

func BindParams(ctx *gin.Context, obj interface{}) (ok bool, res *internal.ValidateReturn) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := ctx.ShouldBindQuery(field); err != nil {
			return false, &internal.ValidateReturn{
				Message: err.Error(),
			}
		}
		res, err := internal.ValidateRequest(&field, "param")
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
