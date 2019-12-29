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
	Uid       uint32    `json:"uid"`
	Username  string    `json:"username"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Birthday  int64     `json:"birthday"`
	Website   string    `json:"website"`
	Phone     string    `json:"phone"`
	AvatarUrl string    `json:"avatar_url"`
}

func GetUserInfo(uid uint32) (interface{}, error) {
	user, err := exitUserByUID(uid)
	if err != nil {
		return nil, err
	}
	tags := mustGetTagsForUser(uid)
	categories := mustGetCategoryForUser(uid)
	result := map[string]interface{}{
		"user":       user,
		"tags":       tags,
		"categories": categories,
	}
	return result, nil
}

// internal method
func getUserInfo(uid uint32) (*User, error) {
	return exitUserByUID(uid)
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
	err := db.Exec("select * from user where uid = ?", uid).Find(&user).Error
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
		Birthday:  data["birthday"].(int64),
		Website:   data["website"].(string),
	}
	return user
}
