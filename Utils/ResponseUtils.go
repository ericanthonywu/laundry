package Utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Constant/APIResponse"
	"laundry/Model"
	"net/http"
)

func errorResponse(err error, context string, httpStatus int) error {
	message := context
	if httpStatus >= 500 {
		message = "failed to " + message
		fmt.Println(message+": ", err)
	}

	return echo.NewHTTPError(httpStatus, message)
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
	return errorResponse(nil, "bad request", http.StatusBadRequest)
}

func OkResponseMessage(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Model.NewDefaultResponse(message, data))
}

func OkResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Model.NewDefaultResponse(APIResponse.DefaultMessageResponse, data))
}
