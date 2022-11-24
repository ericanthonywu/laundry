package APIResponse

const (
	RequestOtpSuccess    = "Success! Please verify your phone number by entering the otp code sent to your number"
	RequestOtpCanRetryAt = "Please wait before retrying to sent sms"
	VerifyOtpBadRequest  = "Oops! Your otp code is not found or has been expired"
	VerifyOtpWrongOtp    = "Oops! Your otp is wrong"
	VerifyOtpSuccess     = "Phone number verified!"
)
