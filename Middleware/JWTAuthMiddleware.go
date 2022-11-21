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

func getToken(next echo.HandlerFunc, role string) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("token")

		if tokenString == "" {
			return c.JSON(http.StatusInternalServerError, Model.ErrorResponse{Message: "token is required"})
		}

		err, secretToken, identifier := Utils.GenerateSecretTokenAndIdentifier(c, role)

		if err != nil {
			return err
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return c.JSON(http.StatusInternalServerError, Model.ErrorResponse{Message: "failed to parse jwt, check log"}), fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretToken), nil
		})

		if err != nil {
			return Utils.JWTErrorResponse(err, c)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return Utils.JWTErrorResponse(err, c)
		}

		jwtRole := claims["role"].(string)

		if jwtRole != identifier {
			return c.JSON(http.StatusBadRequest, Model.ErrorResponse{Message: "role invalid"})
		}

		return next(c)
	}
}
