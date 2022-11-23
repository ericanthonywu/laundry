package Utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Model"
	"net/http"
)

func errorResponse(err error, context string, httpStatus int) error {
	message := "failed to " + context
	if httpStatus >= 500 {
		fmt.Println(message+": ", err)
	}

	return echo.NewHTTPError(httpStatus, Model.NewErrorResponse(message, err))
}

func JWTErrorResponse(err error) error {
	return errorResponse(err, "parse jwt", http.StatusInternalServerError)
}

func DBErrorResponse(err error) error {
	return errorResponse(err, "execute query", http.StatusInternalServerError)
}

func BadRequestResponse(message string) error {
	return errorResponse(nil, message, http.StatusBadRequest)
}