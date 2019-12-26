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

package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"outofmemory/errors"
	"outofmemory/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = errors.CodeNoError
		token := c.GetHeader("Authorization")
		if token == "" {
			code = errors.CodeUnauthorized
		} else {
			_, err := utils.CheckToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errors.CodeTokenTimeout
				default:
					code = errors.CodeCheckTokenFailed
				}
			}
		}
		if code != errors.CodeNoError {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": errors.GetMsg(code),
				"data": make([]interface{}, 0),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
