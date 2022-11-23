package Utils

import (
	"laundry/Config"
	"laundry/Model/Database"
)

func CheckUserPhoneNumberExists(phoneNumber string) (bool, error) {
	var userExists bool

	err := Config.Db().
		Model(&Database.User{}).
		Select("1").
		Where("phone_number", phoneNumber).
		Find(&userExists).Error

	return userExists, err
}