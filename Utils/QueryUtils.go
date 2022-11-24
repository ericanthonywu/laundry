package Utils

import (
	"laundry/Lib"
	"laundry/Model/Database"
)

func CheckUserPhoneNumberExists(phoneNumber string) (bool, error) {
	var userExists bool

	err := Lib.DB.
		Model(&Database.User{}).
		Select("1").
		Where("phone_number", phoneNumber).
		Find(&userExists).Error

	return userExists, err
}
