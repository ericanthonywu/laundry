package UserService

import (
	"laundry/Lib"
	"laundry/Model"
	"laundry/Model/Database"
)

func GetProfile(id string) Model.UserProfileResponse {
	var user Database.User
	var userProfileResponse = Model.UserProfileResponse{}
	if err := Lib.DB.
		Model(&user).
		Where("id = ?", id).
		Select("id", "name", "phone_number", "address", "email_address").
		Take(&userProfileResponse).
		Error; err != nil {
		panic(err)
	}
	return userProfileResponse
}
