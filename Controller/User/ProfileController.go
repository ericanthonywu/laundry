package User

import (
	"github.com/labstack/echo/v4"
	"laundry/Model"
	"laundry/Model/Database"
	"laundry/Services/UserService"
	"laundry/Utils"
)

func GetProfile(c echo.Context) error {
	id, _ := Utils.GetJwtClaims(c)
	profileResponse := UserService.GetProfile(id)

	return Utils.OkResponse(c, profileResponse)
}

func UpdateProfile(c echo.Context) error {
	request := new(Model.UpdateProfileRequest)

	if err := c.Bind(request); err != nil {
		return Utils.BadRequestResponse("failed to bind request")
	}

	id, _ := Utils.GetJwtClaims(c)

	user := &Database.User{
		Name:         request.Name,
		PhoneNumber:  request.PhoneNumber,
		EmailAddress: request.EmailAddress,
		Address:      request.Address,
	}
	user.ID = id

	UserService.UpdateProfile(user)

	return Utils.OkResponse(c, nil)
}
