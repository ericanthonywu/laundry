package User

import (
	"github.com/labstack/echo/v4"
	"laundry/Constant/APIResponse"
	"laundry/Model"
	"laundry/Services/UserService"
	"laundry/Utils"
)

func RequestOTP(c echo.Context) error {
	request := new(Model.UserRequestOTPRequest)

	if err := c.Bind(request); err != nil {
		return Utils.BadRequestResponse("failed to bind request")
	}

	if !request.IsValid() {
		return Utils.BadRequestResponse("request invalid")
	}

	expiresAt, err := UserService.RequestOtpService(request)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, APIResponse.RequestOtpSuccess, Model.UserRequestOTPResponse{
		ExpireAt: expiresAt,
	})
}

func VerifyOTP(c echo.Context) error {
	request := new(Model.UserVerifyOTPRequest)

	if err := c.Bind(request); err != nil {
		return Utils.BadRequestResponse("failed to bind request")
	}

	if !request.IsValid() {
		return Utils.BadRequestResponse("bad request")
	}

	token, err := UserService.VerifyOTP(request)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, APIResponse.VerifyOtpSuccess, Model.UserVerifyOTPResponse{Token: token})
}
