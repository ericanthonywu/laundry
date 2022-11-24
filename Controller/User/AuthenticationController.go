package User

import (
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"laundry/Model"
	"laundry/Services/UserService"
	"laundry/Utils"
	"net/http"
)

func RequestOTP(c echo.Context) error {
	request := new(Model.UserRequestOTPRequest)

	if err := c.Bind(request); err != nil {
		return Utils.BadRequestResponse("failed to bind request")
	}

	if !request.IsValid() {
		return Utils.BadRequestResponse("request invalid")
	}

	if err := UserService.RequestOtpService(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		Model.NewDefaultResponse(
			Constant.RequestOtpSuccess,
			nil,
		),
	)
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

	return c.JSON(http.StatusOK,
		Model.NewDefaultResponse(
			Constant.VerifyOtpSuccess,
			Model.UserVerifyOTPResponse{Token: token},
		),
	)
}
