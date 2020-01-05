package service

import "outofmemory/models"

type UserInfo struct {
	UID      uint32
	Gender   string
	Email    string
	NickName string
	Avatar   string
	BirthDayYear uint32
	BirthDayMonth uint8
	BirthDayDay uint8
	WebSite  string
	Phone    string
}

func (userInfo *UserInfo) UpdateUserInfo() error {
	data := makeUserData(userInfo)
	return models.UpdateUserInfo(data)
}

func (userInfo *UserInfo) GetUserInfo() (interface{}, error) {
	return models.GetUserInfo(userInfo.UID)
}

func makeUserData(user *UserInfo) map[string]interface{} {
	data := map[string]interface{}{
		"uid":        user.UID,
		"gender":     user.Gender,
		"email":      user.Email,
		"phone":      user.Phone,
		"website":    user.WebSite,
		"birthday_y":   user.BirthDayYear,
		"birthday_mo":   user.BirthDayMonth,
		"birthday_d":   user.BirthDayDay,
		"avatar_url": user.Avatar,
	}
	return data
}
