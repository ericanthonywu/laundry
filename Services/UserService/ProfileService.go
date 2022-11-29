package UserService

import (
	"laundry/Lib"
	"laundry/Model"
	"laundry/Model/Database"
)

func GetProfile(id uint) Model.UserProfileResponse {
	var user Database.User
	user.ID = id
	var userProfileResponse = Model.UserProfileResponse{}
	if err := Lib.DB.
		Model(&user).
		Where(&user).
		Select("name", "phone_number", "address", "email_address").
		Take(&userProfileResponse).
		Error; err != nil {
		panic(err)
	}
	userProfileResponse.ID = id
	return userProfileResponse
}

func UpdateProfile(user *Database.User) {
	if err := Lib.DB.Model(&user).Updates(user).Where("id = ?", user.ID); err != nil {
		panic(err)
	}
}
