package User

import (
	"github.com/labstack/echo/v4"
	"laundry/Constant/APIResponse"
	"laundry/Services/UserService"
	"laundry/Utils"
)

func GetProfile(c echo.Context) error {
	id, _ := Utils.GetJwtClaims(c)
	profileResponse := UserService.GetProfile(id)

	return Utils.OkResponse(c, APIResponse.GetProfile, profileResponse)
}
