package Utils

import (
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"laundry/Model"
	"net/http"
	"os"
)

func GenerateSecretTokenAndIdentifier(c echo.Context, role string) (error, string, string) {
	secretToken := ""
	identifier := ""

	switch role {
	case Constant.User:
		secretToken = os.Getenv("JWTUSERSECRETTOKEN")
		identifier = os.Getenv("JWTUSERIDENTIFIER")
	case Constant.Admin:
		secretToken = os.Getenv("JWTADMINSECRETTOKEN")
		identifier = os.Getenv("JWTADMINIDENTIFIER")
	default:
		return c.JSON(http.StatusBadRequest, Model.ErrorResponse{Message: "role not found"}), "", ""
	}

	if secretToken == "" || identifier == "" {
		return c.JSON(http.StatusBadRequest, Model.ErrorResponse{Message: "secret token or identifier is empty"}), "", ""
	}

	return nil, secretToken, identifier
}
