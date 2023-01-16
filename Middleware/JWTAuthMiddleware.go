package Middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"laundry/Model"
	"laundry/Utils"
	"net/http"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return getToken(next, Constant.User)
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return getToken(next, Constant.Admin)
}

func getToken(next echo.HandlerFunc, role string) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("token")

		if tokenString == "" {
			return c.JSON(http.StatusInternalServerError, Model.ErrorResponse{Message: "token is required"})
		}

		err, secretToken, identifier := Utils.GenerateSecretTokenAndIdentifier(role)

		if err != nil {
			return err
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return Utils.JWTErrorResponse(err), fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretToken), nil
		})

		if err != nil {
			return Utils.JWTErrorResponse(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return Utils.JWTErrorResponse(err)
		}

		jwtRole := claims[Constant.UserClaimsRole].(string)

		Utils.SetJwtClaims(c, claims[Constant.UserClaimsId].(string), jwtRole)

		if jwtRole != identifier {
			return Utils.BadRequestResponseWithMessage("role invalid")
		}

		return next(c)
	}
}
