package util

import (
	"context"
	"go-architecture/dto/apperror"
	"go-architecture/pkg/ctxutil"
	"net/http"
	"reflect"
	"unsafe"
)

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Must(err error) bool {
	if err != nil {
		panic(err)
	}
	return false
}

func BadRequestError(ctx context.Context, errorCode string) *apperror.AppError {
	lang := ctxutil.GetAcceptLanguageFromCtx(ctx)
	return apperror.NewAppErrorWithStatus(http.StatusBadRequest, errorCode, "", lang, nil)
}