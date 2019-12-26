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
	"github.com/go-redis/redis"
	"log"
	"outofmemory/settings"
)

var Client *redis.Client

func Setup() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     settings.RedisSetting.Host + settings.RedisSetting.Port,
		Password: settings.RedisSetting.PassWord,
		PoolSize: settings.RedisSetting.PoolSize,
	})
	pong, err := Client.Ping().Result()
	if err != nil {
		log.Fatalf("redis client start error : %v", err)
		return err
	}
	log.Printf("pong : %v", pong)
	return nil
}

func GetRedisClient() *redis.Client  {
	return Client
}
