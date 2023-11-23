package common

import (
	"forimoc.com/Heimbot/model/ModelAdmin"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// 自定义的token密钥
var jwtKey = []byte("97c3c9d7-ab09-4bee-8147-b637afdd5f05")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 发布token
// admin实例 => token字符串
func ReleaseToken(admin ModelAdmin.Admin) (string, error) {
	// 获取参数 expirationDay
	expirationDay := viper.GetDuration("jwt.expirationDay")
	// 有效期 expirationDay 天
	expirationTime := time.Now().Add(expirationDay * 24 * time.Hour)
	claims := &Claims{
		UserID: admin.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Heimbot",
			Subject:   "admin token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析token
// token字符串 => token字符串, token的claims细节, err错误
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
