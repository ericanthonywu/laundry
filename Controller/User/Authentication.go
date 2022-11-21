package User

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"laundry/Model"
	"laundry/Utils"
	"net/http"
)

func Login(c echo.Context) error {
	request := new(Model.UserLoginRequest)

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	err, secretToken, identifier := Utils.GenerateSecretTokenAndIdentifier(c, Constant.User)
	if err != nil {
		return err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ""
	claims["id"] = ""
	claims["role"] = identifier

	t, err := token.SignedString([]byte(secretToken))

	if err != nil {
		return Utils.JWTErrorResponse(err, c)
	}

	return c.JSON(http.StatusOK,
		Model.NewDefaultResponse(
			"login success",
			Model.UserLoginResponse{
				Token: t,
				Id:    "",
			},
		),
	)
}
