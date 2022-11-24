package UserService

import (
	"laundry/Lib"
	"laundry/Model"
	"laundry/Model/Database"
	"strconv"
)

func GetProfile(id string) Model.UserProfileResponse {
	var user Database.User
	uintId, _ := strconv.ParseUint(id, 10, 64)
	user.ID = uint(uintId)
	var userProfileResponse = Model.UserProfileResponse{}
	if err := Lib.DB.
		Model(&user).
		Where(&user).
		Select("name", "phone_number", "address", "email_address").
		Take(&userProfileResponse).
		Error; err != nil {
		panic(err)
	}
	userProfileResponse.ID = uint(uintId)
	return userProfileResponse
}
