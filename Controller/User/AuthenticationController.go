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
		return Utils.BadRequestResponseWithMessage("failed to bind request")
	}

	if !request.IsValid() {
		return Utils.BadRequestResponse()
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
		return Utils.BadRequestResponseWithMessage("failed to bind request")
	}

	if !request.IsValid() {
		return Utils.BadRequestResponse()
	}

	token, err := UserService.VerifyOTP(request)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, APIResponse.VerifyOtpSuccess, Model.UserVerifyOTPResponse{Token: token})
}
