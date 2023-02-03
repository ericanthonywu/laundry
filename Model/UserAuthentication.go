package Model

import "time"

type UserRequestOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func (req UserRequestOTPRequest) IsValid() bool {
	return req.PhoneNumber != ""
}

type UserRequestOTPResponse struct {
	ExpireAt time.Time `json:"expire_at"`
}

type UserVerifyOTPRequest struct {
	OtpCode     string `json:"otp_code"`
	PhoneNumber string `json:"phone_number"`
}

func (req UserVerifyOTPRequest) IsValid() bool {
	return req.PhoneNumber != "" || req.OtpCode != "" || len(req.OtpCode) == 6
}

type UserVerifyOTPResponse struct {
	Token string `json:"token"`
}
