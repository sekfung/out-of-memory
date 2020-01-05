// Copyright 2019 sekfung
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"github.com/jinzhu/gorm"
	"outofmemory/errors"
)

type User struct {
	BaseModel
	Uid       uint32 `gorm:"not null;index:uid" json:"uid"`
	Username  string `gorm:"not null;index:username" json:"username"`
	Gender    string `gorm:"not null;default:'m'" json:"gender"`
	Email     string `gorm:"not null;default:''"json:"email"`
	BirthdayY uint32 `gorm:"type:smallint;default:1900" json:"birthday_y"`
	BirthdayM uint8  `gorm:"type:smallint;default:1" json:"birthday_m"`
	BirthdayD uint8  `gorm:"type:smallint;default:1" json:"birthday_d"`
	Website   string `gorm:"not null;default:''"json:"website"`
	Phone     string `gorm:"not null;default:''" json:"phone"`
	AvatarUrl string `gorm:"not null;default:''"json:"avatar_url"`
}

func GetUserInfo(uid uint32) (interface{}, error) {
	user := User{}
	err := db.Where("uid = ?", uid).Find(&user).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrUserNotExist
		default:
			return nil, errors.ErrUnknownError
		}
	}
	return &user, nil
}

func UpdateUserInfo(data map[string]interface{}) error {
	var (
		userData = parseUserData(data)
	)
	user, err := exitUserByUID(userData.Uid)
	if err != nil {
		return err
	}

	err = db.Table("user").Where("uid = ?", user.Uid).Update(&userData).Error
	return err
}

func exitUserByUID(uid uint32) (*User, error) {
	user := User{}
	err := db.Select("uid").Where("uid = ?", uid).Find(&user).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrUserNotExist
		default:
			return nil, errors.ErrUnknownError
		}
	}
	return &user, nil
}

func parseUserData(data map[string]interface{}) User {
	user := User{
		Uid:       data["uid"].(uint32),
		Gender:    data["gender"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
		AvatarUrl: data["avatar_url"].(string),
		BirthdayY: data["birthday_y"].(uint32),
		BirthdayM: data["birthday_m"].(uint8),
		BirthdayD: data["birthday_d"].(uint8),
		Website:   data["website"].(string),
	}
	return user
}
