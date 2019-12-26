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
	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"outofmemory/errors"
	"time"
)

type UserAuth struct {
	Id             uint32 `json:"-"`
	Uid            uint32 `json:"uid"`             // user id
	Identifier     string `json:"identifier"`      // login identity
	Credential     string `json:"credential"`      // password or token
	IdentifierFrom uint8  `json:"identifier_from"` // identifier from (0: site in , 1: site out)
	IdentityType   string `json:"identity_type"`   // login type (email, phone, username, github, weibo...btw: email, phone, username is belong to site in)
	Verified       bool   `json:"verified"`        // is verified
	VerifyDate     int64  `json:"verify_date"`     // verify date
	LastLoginTime  int64  `json:"last_login_time"`
	CreatedAt      int64  `xorm:"created" json:"-"`
	UpdatedAt      int64  `xorm:"updated" json:"-"`
	DeletedAt      int64  `xorm:"deleted" json:"-"`
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
	userAuth.LastLoginTime = time.Now().Unix()

	// generate bcrypt password, if register by website inside
	if userAuth.IdentifierFrom == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(userAuth.Credential), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.ErrUnknownError
		}
		userAuth.Credential = string(hash)
	}
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if _, err = session.Insert(&userAuth); err != nil {
		session.Rollback()
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
	if _, err := session.Insert(&user); err != nil {
		session.Rollback()
		return &userAuth, errors.ErrUnknownError
	}

	err = session.Commit()

	return &userAuth, nil
}

func CheckUserAuth(param []byte) (*UserAuth, error) {
	var (
		userAuthForm UserAuth
		userResult   UserAuth
	)
	_ = json.Unmarshal(param, &userAuthForm)

	err := engine.SQL("select uid, credential from user_auth where identifier = ? and identity_type = ?",
		userAuthForm.Identifier, userAuthForm.IdentityType).Find(&userResult)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
			return nil, errors.ErrUserNotExist
		default:
			return nil, errors.ErrUnknownError
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Credential), []byte(userAuthForm.Credential))
	if err != nil {
		return nil, errors.ErrIncorrectIdentifierOrCredential
	}

	return &userResult, nil
}

func CheckUserExist(identifier string, identityType string) (bool, error) {
	var userAuth UserAuth
	isExist, err := engine.Select("uid").Where("identifier = ?", identifier).And("identity_type = ?", identityType).Get(&userAuth)
	if !isExist {
		return false, errors.ErrUserNotExist
	}
	//err := engine.SQL("select uid from user_auth where identifier = ? and identity_type = ?", identifier, identityType).Find(&userAuth)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
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
