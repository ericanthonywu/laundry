package Database

import "gorm.io/gorm"

type Laundry struct {
	gorm.Model
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
}

func (Laundry) TableName() string {
	return "laundry"
}
