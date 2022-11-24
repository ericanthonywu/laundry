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

func RequestOtpService(request *Model.UserRequestOTPRequest) error {
	OtpRetrySec := Utils.GetEnvInt("OTP_RETRY_SECONDS")
	OtpExpireSec := Utils.GetEnvInt("OTP_EXPIRE_SECONDS")

	otp := Utils.GenerateOtpCode()
	canRetryAt := time.Now().Local().Add(time.Second * time.Duration(OtpRetrySec))
	expiresAt := time.Now().Local().Add(time.Second * time.Duration(OtpExpireSec))

	Utils.SetUserRedisOtp(request.PhoneNumber, otp)

	if err := Lib.DB.Create(&Database.UserOtpRequest{
		OtpCode:     otp,
		PhoneNumber: request.PhoneNumber,
		CanRetryAt:  canRetryAt,
		ExpiresAt:   expiresAt,
	}).Error; err != nil {
		return Utils.DBErrorResponse(err)
	}

	Lib.SendSMS(request.PhoneNumber,
		fmt.Sprintf(Constant.RequestOTP, otp, OtpExpireSec/60),
	)

	return nil
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
