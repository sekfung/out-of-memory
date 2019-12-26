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
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"outofmemory/settings"
)

var engine *xorm.Engine

func Setup() error {
	err := initDB()
	err = syncDB()
	return err
}

func initDB() error {
	var err error
	databaseSettings := settings.DatabaseSetting
	serverSettings := settings.ServerSetting

	var dbName = databaseSettings.ProName
	if serverSettings.RunMode == "debug" {
		dbName = databaseSettings.DevName
	}
	engine, err = xorm.NewEngine(databaseSettings.Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8",
		databaseSettings.User,
		databaseSettings.Password,
		databaseSettings.Host,
		databaseSettings.Port,
		dbName))
	engine.ShowSQL(serverSettings.RunMode == "debug")
	engine.ShowExecTime(serverSettings.RunMode == "debug")
	if err != nil {
		log.Fatalf("Open engine error : %v", err)
		return err
	}
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	return err
}

func syncDB() error {
	err := engine.Sync2(new(Article), new(Category), new(CategoryArticle), new(Tag), new(TagArticle), new(User), new(UserAuth))
	return err
}

func CloseDB() {
	defer engine.Close()
}
