package UserService

import (
	"laundry/Lib"
	"laundry/Model"
	"laundry/Model/Database"
	"laundry/Utils"
)

func GetProfile(id uint) Model.UpdateProfile {
	var user Database.User
	user.ID = id
	var userProfileResponse = Model.UpdateProfile{}
	if err := Lib.DB.
		Model(&user).
		Where(&user).
		Select("name", "phone_number", "address", "email_address").
		Take(&userProfileResponse).
		Error; err != nil {
		panic(err)
	}

	return userProfileResponse
}

func GetUserCoin(id uint) uint64 {
	var user Database.User
	user.ID = id

	data, existValue := Utils.GetUserCoinRedis(id)

	if !existValue {
		if err := Lib.DB.
			Model(&user).
			Where(&user).
			Select("coin").
			Take(&data).
			Error; err != nil {
			panic(err)
		}
	}

	return data
}

func UpdateProfile(user *Database.User) {
	if err := Lib.DB.Model(&user).Updates(user).Where("id = ?", user.ID).Error; err != nil {
		panic(err)
	}
}
