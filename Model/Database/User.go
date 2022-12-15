package Database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string `gorm:"default:null"`
	PhoneNumber  string
	EmailAddress string    `gorm:"default:null"`
	Address      string    `gorm:"default:null"`
	Birthdate    time.Time `gorm:"default:null"`
	gender       string    `gorm:"default:null"`
	LastLoginAt  time.Time `gorm:"default:null"`
	lat          float64   `gorm:"default:null"`
	long         float64   `gorm:"default:null"`
	Coin         uint64    `gorm:"default:0"`
}

func (User) TableName() string {
	return "user"
}
