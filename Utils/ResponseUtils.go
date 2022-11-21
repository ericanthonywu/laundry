package Utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Model"
	"net/http"
)

func errorResponse(err error, c echo.Context, context string) error {
	message := "failed to " + context
	fmt.Println(message+": ", err)
	return c.JSON(http.StatusInternalServerError, Model.ErrorResponse{Message: message, Error: err})
}

func JWTErrorResponse(err error, c echo.Context) error {
	return errorResponse(err, c, "parse jwt")
}
