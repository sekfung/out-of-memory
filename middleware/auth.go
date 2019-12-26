package middleware

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"outofmemory/errors"
	"outofmemory/utils"
	"strconv"
)

func CheckUserPermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		userClaims, err := utils.CheckToken(token)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code": errors.CodeInvalidParam,
				"msg": errors.GetMsg(errors.CodeInvalidParam),
				"data": make([]interface{}, 0),
			})
			context.Abort()
		}
		uid := strconv.FormatUint(uint64(userClaims.UID), 10)
		id := com.StrTo(context.Param("id")).String()
		if uid != id {
			context.JSON(http.StatusForbidden, gin.H{
				"code": errors.CodeForbidden,
				"msg": errors.GetMsg(errors.CodeForbidden),
				"data": make([]interface{}, 0),
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
