package utils

import (
	"github.com/dgrijalva/jwt-go"
	"outofmemory/settings"
	"time"
)

type userClaims struct {
	UID uint32
	jwt.StandardClaims
}

var jwtSecret = []byte(settings.ApiSetting.JwtSecret)

func GenerateToken(uid uint32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour * 365)
	claims := userClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "out-of-memory",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func CheckToken(token string) (*userClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*userClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
