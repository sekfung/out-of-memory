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

package service

import (
	"outofmemory/models"
)

type Auth struct {
	Identifier     string
	Credential     string
	IdentityType   string
	IdentifierFrom uint8
}

func (auth *Auth) DoAuth(param []byte) (*models.UserAuth, error) {
	return models.CheckUserAuth(param)
}

// Register User
func (auth *Auth) Register() (*models.UserAuth, error) {
	data := makeAuthData(auth)
	return models.Register(data)
}

func (auth *Auth) CheckUserExist() (bool, error) {
	return models.CheckUserExist(auth.Identifier, auth.IdentityType)
}

func makeAuthData(auth *Auth) map[string]interface{}  {
	data := map[string]interface{} {
		"identifier": auth.Identifier,
		"identifier_from": auth.IdentifierFrom,
		"credential": auth.Credential,
		"identity_type": auth.IdentityType,
	}
	return data
}
