package Database

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Name string `json:"name"`
}

func (Services) TableName() string {
	return "service"
}
