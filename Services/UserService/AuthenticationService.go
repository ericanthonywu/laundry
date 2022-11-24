package UserService

import (
	"fmt"
	"laundry/Constant"
	"laundry/Constant/APIResponse"
	"laundry/Lib"
	"laundry/Model"
	"laundry/Model/Database"
	"laundry/Utils"
	"time"
)

func RequestOtpService(request *Model.UserRequestOTPRequest) (time.Time, error) {
	OtpExpireSec := Utils.GetEnvInt("OTP_EXPIRE_SECONDS")

	var otp string
	otp = Utils.GetUserRedisOtp(request.PhoneNumber)

	if otp != "" {
		return time.Now(), Utils.BadRequestResponse(APIResponse.RequestOtpCanRetryAt)
	}

	expiresAt := time.Now().Local().Add(time.Second * time.Duration(OtpExpireSec))
	otp = Utils.GenerateOtpCode()

	Utils.SetUserRedisOtp(request.PhoneNumber, otp)

	if err := Lib.DB.Create(&Database.UserOtpRequest{
		OtpCode:     otp,
		PhoneNumber: request.PhoneNumber,
		ExpiresAt:   expiresAt,
	}).Error; err != nil {
		return time.Now(), Utils.DBErrorResponse(err)
	}

	Lib.SendSMS(request.PhoneNumber,
		fmt.Sprintf(Constant.RequestOTP, otp, OtpExpireSec/60),
	)

	return expiresAt, nil
}

func VerifyOTP(request *Model.UserVerifyOTPRequest) (string, error) {
	otp := Utils.GetUserRedisOtp(request.PhoneNumber)

	if otp == "" {
		return "", Utils.BadRequestResponse(APIResponse.VerifyOtpBadRequest)
	}

	if otp != request.OtpCode {
		return "", Utils.BadRequestResponse(APIResponse.VerifyOtpWrongOtp)
	}

	defer Utils.DelUserRedisOtp(request.PhoneNumber)

	var userData = Database.User{PhoneNumber: request.PhoneNumber}

	err := Lib.DB.
		Where(userData).
		Select("id").
		First(&userData).Error

	if Utils.IsDBNotFound(err) && err != nil {
		err = Lib.DB.Create(&userData).Error
	}

	if err != nil {
		panic(err)
	}

	token, err := Utils.GenerateJwtToken(userData.ID, Constant.User)

	if err != nil {
		panic(err)
	}

	return token, nil
}
