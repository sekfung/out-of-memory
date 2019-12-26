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

package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"outofmemory/api"
	"outofmemory/errors"
	"outofmemory/models"
	"outofmemory/service"
	"outofmemory/settings"
	"strconv"
)

type tagForm struct {
	TagID uint32 `form:"id" validate:"omitempty,numeric,min=1" json:"tag_id"`
	State uint8  `form:"state" validate:"omitempty,oneof=0 1" json:"state"`
	Name  string `form:"name" validate:"omitempty,max=20,min=1" json:"name"`
}

// @Summary Get Tags
// @Produce  json
// @Tags Tag
// @Param page query string false "Page number, default 1"
// @Param per_page query string false "Limit how much results returned per page, default 20"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form tagForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	// default query
	pageReq := c.DefaultQuery("page", "1")
	page := com.StrTo(pageReq).MustInt()
	perPageReq := c.DefaultQuery("per_page", strconv.Itoa(settings.ApiSetting.PageSize))
	perPage := com.StrTo(perPageReq).MustInt()

	tagService := service.Tag{
		Page:    page,
		PerPage: perPage,
	}
	tags, err := tagService.GetTags()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
	}
	appG.Response(errors.CodeNoError, tags)
}

// @Summary Get Tag
// @Produce  json
// @Tags Tag
// @Param id query int true "Category Id,"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/tags/:id [get]
func GetTagById(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form tagForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appG.Response(errors.CodeInvalidParam, nil)
		return
	}
	tagService := service.Tag{TagID: uint32(tagID)}
	tag, err := tagService.GetTagByID()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, tag)
}

// @Summary Update Tag
// @Produce  json
// @Tags Tag
// @Param id query int true "Category Id"
// @Param name query string true "Category name"
// @Param state query int false "Category state"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/tags/:id [put]
func UpdateTag(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form tagForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	tagState := c.DefaultQuery("state", "1")
	state := com.StrTo(tagState).MustUint8()
	tagService := service.Tag{
		TagID: uint32(tagID),
		Name:  form.Name,
		State: state,
	}
	tag, err := tagService.UpdateTag()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, tag)
}

// @Summary Add Tag
// @Produce  json
// @Tags Tag
// @Param state query string false "Tag state, 0: disable, 1:enable, default 1"
// @Param name query string true "Tag name, max length 20 characters"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form tagForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	tagService := service.Tag{
		Name:  form.Name,
		State: models.TagStateToUint["enable"],
	}
	tag, err := tagService.AddTag()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, tag)
}

// @Summary Delete Tag
// @Produce  json
// @Tags Tag
// @Param id query int true "Category Id"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/tags [delete]
func DeleteTagByID(c *gin.Context)  {
	var (
		appG = api.Gin{C: c}
		form tagForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	tagService := service.Tag{
		TagID: uint32(tagID),
	}
	err = tagService.DeleteTagById()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, nil)
}
