package Utils

import (
	"laundry/Constant"
	"laundry/Lib"
	"time"
)

func generateUserOtpKey(phoneNumber string) string {
	return Constant.UserOtpRedisKey + "-" + phoneNumber
}

func SetUserRedisOtp(phoneNumber string, otp string) {
	OtpExpireSec := GetEnvInt("OTP_EXPIRE_SECONDS")
	Lib.RDBSet(generateUserOtpKey(phoneNumber), otp, time.Duration(OtpExpireSec)*time.Second)
}

func GetUserRedisOtp(phoneNumber string) string {
	return Lib.RDBGet(generateUserOtpKey(phoneNumber))
}

func DelUserRedisOtp(phoneNumber string) {
	Lib.RDBDel(generateUserOtpKey(phoneNumber))
}
