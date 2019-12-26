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

package api

import (
	"github.com/gin-gonic/gin"
	"outofmemory/errors"
	"outofmemory/service"
	"outofmemory/utils"
)

type registerForm struct {
	Identifier     string `form:"identifier" validate:"required" json:"identifier"`
	Credential     string `form:"credential" validate:"required" json:"credential"`
	IdentityType   string `form:"identity_type" validate:"required" json:"identity_type"`
	IdentifierFrom uint8  `form:"identifier_from" validate:"oneof=0 1" json:"identifier_from"`
}

// @Summary Fresh Man Register
// @Produce  json
// @Tags Login
// @Param identifier query string true "User identifier, such as username, email, phone, or uid return by website supported oauth2.0"
// @Param credential query string true "Credential, if user sign in website inside (identifier_from is 0), credential is password, otherwise it's access token"
// @Param identity_type query string true "IdentityType, such as username, email, phone, github, weibo, wechat..."
// @Param identifier_from query int true "IdentifierFrom, range is 0 to 1,  0 means website inside, 1 is outside"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /register [post]
func Register(c *gin.Context) {
	var (
		appG = Gin{C: c}
		form registerForm
	)
	err := BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}

	authService := service.Auth{
		Identifier:     form.Identifier,
		IdentifierFrom: form.IdentifierFrom,
		IdentityType:   form.IdentityType,
		Credential:     form.Credential,
	}
	isExist, err := authService.CheckUserExist()
	// user is existed, no need to repeat registration
	if isExist {
		appG.Response(errors.CodeUserIsExist, nil)
		return
	}
	if err != errors.ErrUserNotExist {
		appG.Response(errors.CodeInternalServerError, nil)
		return
	}
	userAuth, err := authService.Register()
	if err != nil {
		appG.Response(errors.CodeInternalServerError, nil)
		return
	}
	token, err := utils.GenerateToken(userAuth.Uid)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	userService := service.UserInfo{
		UID: userAuth.Uid,
	}
	userInfo, err := userService.GetUserInfo()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, gin.H{
		"token": token,
		"user":  userInfo,
	})
}
