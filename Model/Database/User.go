package Database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	PhoneNumber  string
	EmailAddress string
	Address      string
	Birthdate    time.Time
	gender       string
	LastLoginAt  time.Time
	lat          float64
	long         float64
	Coin         uint64
}

func (User) TableName() string {
	return "user"
}
