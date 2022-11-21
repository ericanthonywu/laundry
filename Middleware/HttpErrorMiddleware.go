package Middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Model"
	"net/http"
)

func httpHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Error(report)

	var (
		code    = report.Code
		message = report.Message
	)

	switch report.Message {
	case "not found":
		code = http.StatusNotFound
		break
	case "sql: no rows in result set":
		code = http.StatusBadRequest
		message = "data not found"
		break
	default:
		code = http.StatusInternalServerError
	}

	if c.JSON(code, Model.ErrorResponse{Message: message}) != nil {
		fmt.Println(err)
	}
}
