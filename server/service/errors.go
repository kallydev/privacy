package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	OK                    = errors.New("ok")
	UnknownError          = errors.New("unknown error")
	InvalidParameterError = errors.New("invalid parameter error")
)

var errorMap = map[error]int{
	OK:                    0,
	UnknownError:          10001,
	InvalidParameterError: 10002,
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(ctx echo.Context, err error, data interface{}) error {
	if err == nil {
		err = OK
	}
	code, ok := errorMap[err]
	if !ok {
		err = UnknownError
		code = errorMap[err]
	}
	return ctx.JSONPretty(http.StatusOK, &Response{
		Code:    code,
		Message: strings.ToLower(err.Error()),
		Data:    data,
	}, "\x20\x20")
}
