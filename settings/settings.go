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

package settings

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host     string
	Port     string
	PassWord string
	PoolSize int
}

var RedisSetting = &Redis{}

type Api struct {
	PageSize  int
	JwtSecret string
}

var ApiSetting = &Api{}

var cfg *ini.File

func LoadSettings() error {
	var err error
	var workspace string
	if workspace, err = os.Getwd(); err != nil {
		log.Fatalln("Fail to get workspace path")
	}
	cfg, err = ini.Load(filepath.Join(workspace, "/conf/app.ini"))
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v", err)
	}
	mapTo("Server", ServerSetting)
	mapTo("Database", DatabaseSetting)
	mapTo("Redis", RedisSetting)
	mapTo("Api", ApiSetting)

	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second
	return err
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo err: %v", err)
	}
}
