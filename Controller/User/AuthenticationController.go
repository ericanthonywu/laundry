package User

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Config"
	"laundry/Constant"
	"laundry/Model"
	"laundry/Model/Database"
	"laundry/Utils"
	"net/http"
	"time"
)

func RequestOTP(c echo.Context) error {
	request := new(Model.UserRequestOTPRequest)

	if err := c.Bind(request); err != nil {
		return Utils.BadRequestResponse("failed to bind request")
	}

	OtpRetrySec := Utils.GetEnvInt("OTP_RETRY_SECONDS")
	OtpExpireSec := Utils.GetEnvInt("OTP_EXPIRE_SECONDS")

	otp := Utils.GenerateOtpCode()
	canRetryAt := time.Now().Local().Add(time.Second * time.Duration(OtpRetrySec))
	expiresAt := time.Now().Local().Add(time.Second * time.Duration(OtpExpireSec))

	otpRequest := Database.UserOtpRequest{
		OtpCode:     otp,
		PhoneNumber: request.PhoneNumber,
		CanRetryAt:  canRetryAt,
		ExpiresAt:   expiresAt,
	}

	if err := Config.Db().Create(&otpRequest).Error; err != nil {
		return Utils.DBErrorResponse(err)
	}

	Config.SendSMS(request.PhoneNumber,
		fmt.Sprintf(Constant.RequestOTP, otp, OtpExpireSec/60),
	)

	return c.JSON(http.StatusOK,
		Model.NewDefaultResponse(
			"otp has been sent",
			nil,
		),
	)
}
