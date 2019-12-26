package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"outofmemory/api"
	"outofmemory/errors"
	"outofmemory/service"
	"outofmemory/settings"
	"strconv"
)

type categoryForm struct {
	CategoryID uint32 `form:"id" validata:"omitempty,numberic,min=1" json:"category_id"`
	Name       string `form:"name" validate:"omitempty,max=20,min=1,max=50" json:"name"`
	State      uint8  `form:"state" validate:"omitempty,oneof=0 1" json:"state"`
}

// @Summary Add Category
// @Produce  json
// @Tags Category
// @Param Name query string false "Category name, limit up to 50 characters"
// @Param State query string false "Category state, 0: disable, 1: enable, default: 1"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/category [post]
func AddCategory(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form categoryForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	categoryService := service.Category{
		Name:  form.Name,
		State: form.State,
	}
	category, err := categoryService.AddCategory()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, category)
}

// @Summary Delete Category
// @Produce  json
// @Tags Category
// @Param Name query string false "Category name, limit up to 50 characters"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/category [post]
func DeleteCategory(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form categoryForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	categoryService := service.Category{
		CategoryID: uint32(categoryID),
	}
	err = categoryService.DeleteCategoryByID()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, nil)
}

func UpdateCategory(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form categoryForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	categoryState := c.DefaultQuery("state", "1")
	state := com.StrTo(categoryState).MustUint8()
	categoryService := service.Category{
		CategoryID: uint32(categoryID),
		Name:       form.Name,
		State:      state,
	}
	category, err := categoryService.UpdateCategory()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, category)
}

func GetCategoryByID(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form categoryForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appG.Response(errors.CodeInvalidParam, nil)
		return
	}
	categoryService := service.Category{CategoryID: uint32(categoryID)}
	category, err := categoryService.GetCategoryByID()
	if err != nil {
		appG.Response(errors.CodeGetTagFailed, nil)
		return
	}
	appG.Response(errors.CodeNoError, category)
}

func GetCategories(c *gin.Context) {
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
	categoryState := c.DefaultQuery("state", "1")
	state := com.StrTo(categoryState).MustUint8()
	tagService := service.Category{
		State:   state,
		Page:    page,
		PerPage: perPage,
	}
	categories, err := tagService.GetCategories()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
	}
	appG.Response(errors.CodeNoError, categories)
}
