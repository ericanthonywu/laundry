package Utils

import (
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
		return BadRequestResponseWithMessage("role not found"), "", ""
	}

	if secretToken == "" || identifier == "" {
		return BadRequestResponseWithMessage("secret token or identifier is empty"), "", ""
	}

	return nil, secretToken, identifier
}

func SetJwtClaims(c echo.Context, id string, role string) {
	c.Set(Constant.UserClaimsId, id)
	c.Set(Constant.UserClaimsRole, role)
}

func GetJwtClaims(c echo.Context) (uint, string) {
	return c.Get(Constant.UserClaimsId).(uint), c.Get(Constant.UserClaimsRole).(string)
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
		return "", err
	}
	return t, nil
}
