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

package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"outofmemory/api/v1"
	_ "outofmemory/docs"
	auth "outofmemory/middleware"
	"outofmemory/settings"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(settings.ServerSetting.RunMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/v1/auth", v1.GetToken)
	r.POST("/v1/register", v1.Register)

	apiV1 := r.Group("/api/v1")
	// tag
	apiV1.GET("/tags", v1.GetTags)
	apiV1.GET("/tags/:id", v1.GetTagById)
	// category
	apiV1.GET("/categories", v1.GetCategories)
	apiV1.GET("/categories/:id",v1.GetCategoryByID)
	// need login
	apiV1.Use(auth.JWT())
	{
		// tag router
		apiV1.PUT("/tags/:id", v1.UpdateTag)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTagByID)

		// category router
		apiV1.PUT("/categories/:id", v1.UpdateCategory)
		apiV1.POST("/categories", v1.AddCategory)
		apiV1.DELETE("/categories/:id",v1.DeleteCategory)

		// article router
		apiV1.POST("/articles", v1.AddArticle)

	}

	// need user permission
	needPermission := r.Group("/api/v1")
	needPermission.Use(auth.JWT())
	needPermission.Use(auth.CheckUserPermission())
	{
		needPermission.PUT("/users/:id", v1.UpdateUserInfo)
		needPermission.GET("/users/:id", v1.GetUserInfo)
	}

	return r
}
