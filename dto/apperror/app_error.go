package apperror

import (
	"fmt"
	"go-architecture/pkg/constant"
	"net/http"
)

type AppError struct {
	Status       int
	ErrorCode    string
	ErrorMessage string
	Lang         string
	ErrorParams  map[string]string
	ErrorFields  []AppErrorField
}

type AppErrorField struct {
	Field        string
	ErrorCode    string
	ErrorMessage string
	ErrorParams  map[string]string
}

func (r *AppError) Error() string {
	return fmt.Sprintf("%s: %s", r.ErrorCode, r.ErrorMessage)
}

func NewAppError(errorCode string, errorMsg string, lang string, errorParams map[string]string) *AppError {
	if len(lang) == 0 {
		lang = constant.DefaultLang
	}

	if errorParams == nil {
		errorParams = make(map[string]string)
	}

	return &AppError{
		Status:       http.StatusInternalServerError,
		ErrorCode:    errorCode,
		ErrorMessage: errorMsg,
		Lang:         lang,
		ErrorParams:  errorParams,
	}
}

func NewAppErrorWithStatus(status int, errorCode string, errorMsg string, lang string, errorParams map[string]string) *AppError {
	if len(lang) == 0 {
		lang = constant.DefaultLang
	}

	if errorParams == nil {
		errorParams = make(map[string]string)
	}

	return &AppError{
		Status:       status,
		ErrorCode:    errorCode,
		ErrorMessage: errorMsg,
		Lang:         lang,
		ErrorParams:  errorParams,
	}
}

