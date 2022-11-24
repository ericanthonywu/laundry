package Model

import (
	"gorm.io/gorm"
)

type UserRequestOTPRequest struct {
	gorm.DB
	PhoneNumber string `json:"phone_number"`
}

func (req UserRequestOTPRequest) IsValid() bool {
	return req.PhoneNumber != ""
}

type UserVerifyOTPRequest struct {
	OtpCode     string `json:"otp_code"`
	PhoneNumber string `json:"phone_number"`
}

func (req UserVerifyOTPRequest) IsValid() bool {
	return req.PhoneNumber != "" || req.OtpCode != ""
}

type UserVerifyOTPResponse struct {
	Token string `json:"token"`
}
