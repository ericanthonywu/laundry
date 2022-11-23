package Database

import (
	"gorm.io/gorm"
	"time"
)

type UserOtpRequest struct {
	gorm.Model
	PhoneNumber string `gorm:"size:20;not null"`
	ExpiresAt   time.Time
	OtpCode     string `gorm:"not null;unique;type:char(6)"`
	CanRetryAt  time.Time
}

func (UserOtpRequest) TableName() string {
	return "user_otp_request"
}
