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

func RequestOtpService(request *Model.UserRequestOTPRequest) (time.Time, time.Time, error) {
	OtpRetrySec := Utils.GetEnvInt("OTP_RETRY_SECONDS")
	OtpExpireSec := Utils.GetEnvInt("OTP_EXPIRE_SECONDS")

	var otp string
	otp = Utils.GetUserRedisOtp(request.PhoneNumber)

	canRetryAt := time.Now().Local().Add(time.Second * time.Duration(OtpRetrySec))
	expiresAt := time.Now().Local().Add(time.Second * time.Duration(OtpExpireSec))
	if otp == "" {
		otp = Utils.GenerateOtpCode()

		Utils.SetUserRedisOtp(request.PhoneNumber, otp)

		if err := Lib.DB.Create(&Database.UserOtpRequest{
			OtpCode:     otp,
			PhoneNumber: request.PhoneNumber,
			CanRetryAt:  canRetryAt,
			ExpiresAt:   expiresAt,
		}).Error; err != nil {
			return time.Now(), time.Now(), Utils.DBErrorResponse(err)
		}
	} else {
		var UserOtpRequest Database.UserOtpRequest
		if err := Lib.DB.Model(&UserOtpRequest).
			Select("can_retry_at").
			Take(&UserOtpRequest).Error; err != nil {
			if Utils.IsDBNotFound(err) {
				Utils.DelUserRedisOtp(request.PhoneNumber)
			}
			panic(err)
		}
		_, min, sec := time.Now().Clock()
		_, userMin, userSec := UserOtpRequest.CanRetryAt.Clock()

		fmt.Println(min, sec, userMin, userSec)
		if time.Now().After(UserOtpRequest.CanRetryAt) {
			return time.Now(), time.Now(), Utils.BadRequestResponse(APIResponse.RequestOtpCanRetryAt)
		}
	}

	Lib.SendSMS(request.PhoneNumber,
		fmt.Sprintf(Constant.RequestOTP, otp, OtpExpireSec/60),
	)

	return canRetryAt, expiresAt, nil
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
