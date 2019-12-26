package v1

import (
	"github.com/gin-gonic/gin"
	"outofmemory/api"
	"outofmemory/errors"
	"outofmemory/service"
	"strconv"
)

type userInfoForm struct {
	UID       uint32 `json:"uid"`
	Email     string `form:"email" validate:"omitempty,email" json:"email,omitempty"`
	NickName  string `form:"nickname" validate:"omitempty" json:"nick_name,omitempty"`
	AvatarURL string `form:"avatar_url" validate:"omitempty,url" json:"avatar_url,omitempty"`
	Gender    string `form:"gender" validate:"omitempty,oneof=f m" json:"gender"`
	WebSite   string `form:"website" validate:"omitempty,url" json:"website"`
}

func GetUserInfo(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
	)
	reqID := c.Param("id")
	uid, err := strconv.ParseUint(reqID, 10, 32)
	if err != nil {
		appG.Response(errors.CodeInvalidParam, nil)
		return
	}
	userService := service.UserInfo{UID: uint32(uid)}
	userInfo, err := userService.GetUserInfo()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, userInfo)
		return
	}
	appG.Response(errors.CodeNoError, userInfo)
}

func UpdateUserInfo(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form userInfoForm
	)
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appG.Response(errors.CodeInvalidParam, nil)
		return
	}
	err = api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	userService := service.UserInfo{
		UID:      uint32(uid),
		Email:    form.Email,
		NickName: form.NickName,
		Avatar:   form.AvatarURL,
		Gender:   form.Gender,
		WebSite:  form.WebSite,
	}
	err = userService.UpdateUserInfo()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, nil)
}
