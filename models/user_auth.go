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
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"outofmemory/errors"
	"time"
)

type UserAuth struct {
	BaseModel
	Uid            uint32     `gorm:"not null;index:uid"json:"uid"`                                     // user id
	Identifier     string     `gorm:"not null;index:identifier"json:"identifier"`                       // login identity
	Credential     string     `gorm:"not null;index:credential"json:"credential"`                       // password or token
	IdentifierFrom uint8      `gorm:"not null;default: 0;index:identifier_from" json:"identifier_from"` // identifier from (0: site in , 1: site out)
	IdentityType   string     `gorm:"not null;index:identity_type"json:"identity_type"`                 // login type (email, phone, username, github, weibo...btw: email, phone, username is belong to site in)
	Verified       bool       `json:"verified"`                                                         // is verified
	VerifyDate     *time.Time `json:"verify_date"`                                                      // verify date
	LastLoginTime  time.Time  `json:"last_login_time"`
}

// 新用户注册
func Register(data map[string]interface{}) (*UserAuth, error) {
	userAuth := parseAuthData(data)
	isExist, checkExistErr := CheckUserExist(userAuth.Identifier, userAuth.IdentityType)
	if isExist || checkExistErr != errors.ErrUserNotExist {
		return nil, checkExistErr
	}

	// generate a unique short user id
	uid := uuid.New().ID()
	userAuth.Uid = uid
	// update login time
	userAuth.LastLoginTime = time.Now()

	// generate bcrypt password, if register by website inside
	if userAuth.IdentifierFrom == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(userAuth.Credential), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.ErrUnknownError
		}
		userAuth.Credential = string(hash)
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&userAuth).Error; err != nil {
		tx.Rollback()
		return &userAuth, errors.ErrUnknownError
	}

	user := User{
		Uid: uid,
	}
	// register by username
	if userAuth.IdentityType == "username" {
		user.Username = userAuth.Identifier
	}
	// register by phone
	if userAuth.IdentityType == "phone" {
		user.Phone = userAuth.Identifier
	}
	// register by email
	if userAuth.IdentityType == "email" {
		user.Email = userAuth.Identifier
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return &userAuth, errors.ErrUnknownError
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return &userAuth, errors.ErrUnknownError
	}

	return &userAuth, nil
}

func CheckUserAuth(param []byte) (*UserAuth, error) {
	var (
		userAuthForm UserAuth
		userResult   UserAuth
	)
	_ = json.Unmarshal(param, &userAuthForm)

	err := db.Select("identifier, credential, identifier_from").
		Where("identifier = ? AND identity_type = ? AND identifier_from = ?", userAuthForm.Identifier, userAuthForm.IdentityType, userAuthForm.IdentifierFrom).
		First(&userResult).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrUserNotExist
		default:
			return nil, errors.ErrUnknownError
		}
	}
	if userResult.IdentifierFrom == 0 {
		err = bcrypt.CompareHashAndPassword([]byte(userResult.Credential), []byte(userAuthForm.Credential))
		if err != nil {
			return nil, errors.ErrIncorrectIdentifierOrCredential
		}
	} else {
		if userAuthForm.Credential == userResult.Credential {
			return &userResult, nil
		} else {
			return nil, errors.ErrIncorrectIdentifierOrCredential
		}
	}
	return &userResult, nil
}

func CheckUserExist(identifier string, identityType string) (bool, error) {
	var userAuth UserAuth
	err := db.Select("uid").Where("identifier = ?", identifier).Where("identity_type = ?", identityType).First(&userAuth).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return false, errors.ErrUserNotExist
		default:
			return false, errors.ErrUnknownError
		}
	}
	return true, nil
}

func parseAuthData(data map[string]interface{}) UserAuth {
	userAuth := UserAuth{
		IdentityType:   data["identity_type"].(string),
		IdentifierFrom: data["identifier_from"].(uint8),
		Credential:     data["credential"].(string),
		Identifier:     data["identifier"].(string),
	}
	return userAuth
}
