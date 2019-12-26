package utils

import "github.com/gin-gonic/gin"

func MustGetUIDFromHeader(c *gin.Context) uint32 {
	token := c.GetHeader("Authorization")
	userClaims, _ := CheckToken(token)
	return userClaims.UID
}