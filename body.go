package bindator

import (
	"io"
	"reflect"

	"github.com/fanesz/bindator/handler"
	"github.com/fanesz/bindator/internal"
	"github.com/gin-gonic/gin"
)

func BindBody(ctx *gin.Context, req interface{}) (res *internal.ValidateReturn) {
	if err := ctx.ShouldBindJSON(req); err != nil {
		if err == io.EOF {
			return handler.Response(false, "Invalid JSON data: unexpected end of JSON input")
		}
		return &internal.ValidateReturn{
			Ok:      false,
			Message: err.Error(),
		}
	}

	res, err := internal.ValidateRequest(&req, "body")
	if err != nil {
		return handler.Response(false, err.Error())
	}
	if res != nil {
		return res
	}

	return handler.Response(true, "")
}

func BindBodies(ctx *gin.Context, obj interface{}) (res *internal.ValidateReturn) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := ctx.ShouldBindJSON(field); err != nil {
			if err == io.EOF {
				return handler.Response(false, "Invalid JSON data: unexpected end of JSON input")
			}
			return handler.Response(false, err.Error())
		}
		res, err := internal.ValidateRequest(&field, "body")
		if err != nil {
			return handler.Response(false, err.Error())
		}
		if res != nil {
			return res
		}
	}

	return handler.Response(true, "")
}
