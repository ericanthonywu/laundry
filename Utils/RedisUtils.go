package Utils

import (
	"laundry/Constant"
	"laundry/Lib"
	"strconv"
	"time"
)

func generateUserOtpKey(phoneNumber string) string {
	return Constant.UserOtpRedisKey + "-" + phoneNumber
}

func generateUserCoinKey(id uint) string {
	return Constant.UserCoinRedisKey + "-" + strconv.Itoa(int(id))
}

func SetUserRedisOtp(phoneNumber string, otp string) {
	OtpExpireSec := GetEnvInt("OTP_EXPIRE_SECONDS")
	Lib.RDBSet(generateUserOtpKey(phoneNumber), otp, time.Duration(OtpExpireSec)*time.Second)
}

func GetUserRedisOtp(phoneNumber string) (string, bool) {
	return Lib.RDBGet(generateUserOtpKey(phoneNumber))
}

func DelUserRedisOtp(phoneNumber string) {
	Lib.RDBDel(generateUserOtpKey(phoneNumber))
}

func SetUserCoinRedis(coin uint, id uint) {
	Lib.RDBSet(generateUserCoinKey(id), coin, 0)
}

func GetUserCoinRedis(id uint) (uint64, bool) {
	return Lib.RDBGetUint(generateUserCoinKey(id))
}
