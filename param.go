package bindator

import (
	"reflect"

	"github.com/fanesz/bindator/handler"
	"github.com/fanesz/bindator/internal"
	"github.com/gin-gonic/gin"
)

func BindParam(ctx *gin.Context, req interface{}) (res *internal.ValidateReturn) {
	if err := ctx.ShouldBindQuery(req); err != nil {
		return handler.Response(false, err.Error())
	}

	res, err := internal.ValidateRequest(&req, "param")
	if err != nil {
		return handler.Response(false, err.Error())
	}
	if res != nil {
		return res
	}

	return handler.Response(true, "")
}

func BindParams(ctx *gin.Context, obj interface{}) (res *internal.ValidateReturn) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := ctx.ShouldBindQuery(field); err != nil {
			return handler.Response(false, err.Error())
		}
		res, err := internal.ValidateRequest(&field, "param")
		if err != nil {
			return handler.Response(false, err.Error())
		}
		if res != nil {
			return res
		}
	}

	return handler.Response(true, "")
}
