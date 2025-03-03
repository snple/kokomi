package util

import (
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func Success(data any) (int, Response) {
	return http.StatusOK, Response{
		Code: 0,
		Data: data,
	}
}

func Error(code int, message string) (int, Response) {
	return http.StatusOK, Response{
		Code:    code,
		Message: message,
	}
}

type Page struct {
	Limit   uint32 `json:"limit,omitempty" form:"limit"`
	Offset  uint32 `json:"offset,omitempty" form:"offset"`
	Search  string `json:"search,omitempty" form:"search"`
	OrderBy string `json:"order_by,omitempty" form:"order_by"`
	Sort    int32  `json:"sort,omitempty" form:"sort"`
}
