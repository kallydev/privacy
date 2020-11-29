package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	OK                    = errors.New("ok")
	InvalidParameterError = errors.New("invalid parameter error")
	QueryFailedError      = errors.New("query failed error")
)

var errorMap = map[error]int{
	OK:                    0,
	InvalidParameterError: 10001,
	QueryFailedError:      10002,
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
		code = 10000
	}
	return ctx.JSONPretty(http.StatusOK, &Response{
		Code:    code,
		Message: strings.ToLower(err.Error()),
		Data:    data,
	}, "\x20\x20")
}
