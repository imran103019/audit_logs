package errorutil

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

// CustomError ...
type CustomError struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

// CustomMultiError ...
type CustomMultiError struct {
	StatusCode int        `json:"status_code"`
	StatusText string     `json:"status_text"`
	Message    url.Values `json:"message"`
}

// NewCustomError returns a new custom error
func NewCustomError(statusCode int, errMsg string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Message:    errMsg,
	}
}

// NewCustomMultiError ...
func NewCustomMultiError(statusCode int, errMsg url.Values) *CustomMultiError {
	return &CustomMultiError{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Message:    errMsg,
	}
}

// SendCustomerIdErrorResponse returns a error
func SendCustomerIdErrorResponse(c echo.Context, statusCode int, errMsg string) error {
	return c.JSON(statusCode, NewCustomError(statusCode, errMsg))
}

// error types
const (
	ErrNoDataFound           = "No data found for the request"
	ErrSomethingWentWrong    = "Something Went Wrong"
)