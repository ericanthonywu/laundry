package Utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Model"
	"net/http"
)

func errorResponse(err error, message string, httpStatus int) error {
	if httpStatus >= 500 {
		message = "failed to " + message
		fmt.Println(message+": ", err)
	}

	return echo.NewHTTPError(httpStatus, http.StatusText(httpStatus))
}

func JWTErrorResponse(err error) error {
	return errorResponse(err, "parse jwt", http.StatusInternalServerError)
}

func DBErrorResponse(err error) error {
	return errorResponse(err, "execute query", http.StatusInternalServerError)
}

func BadRequestResponseWithMessage(message string) error {
	return errorResponse(nil, message, http.StatusBadRequest)
}

func BadRequestResponse() error {
	return BadRequestResponseWithMessage(http.StatusText(http.StatusBadRequest))
}

func OkResponseMessage(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Model.NewDefaultResponse(message, data))
}

func OkResponse(c echo.Context, data interface{}) error {
	return OkResponseMessage(c, http.StatusText(http.StatusOK), data)
}
