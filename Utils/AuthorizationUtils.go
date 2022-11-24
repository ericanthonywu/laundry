package Utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"os"
	"strconv"
)

func GenerateSecretTokenAndIdentifier(role string) (error, string, string) {
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
		return BadRequestResponse("role not found"), "", ""
	}

	if secretToken == "" || identifier == "" {
		return BadRequestResponse("secret token or identifier is empty"), "", ""
	}

	return nil, secretToken, identifier
}

func SetJwtClaims(c echo.Context, id string, role string) {
	c.Set(Constant.UserClaimsId, id)
	c.Set(Constant.UserClaimsRole, role)
}

func GetJwtClaims(c echo.Context) (string, string) {
	userId := fmt.Sprintf("%v", c.Get(Constant.UserClaimsId))
	userRole := fmt.Sprintf("%v", c.Get(Constant.UserClaimsRole))
	return userId, userRole
}

func GenerateJwtToken(id uint, role string) (string, error) {
	err, secretToken, role := GenerateSecretTokenAndIdentifier(role)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims[Constant.UserClaimsId] = strconv.Itoa(int(id))
	claims[Constant.UserClaimsRole] = role

	t, err := token.SignedString([]byte(secretToken))

	if err != nil {
		panic(err)
	}
	return t, nil
}
