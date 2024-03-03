package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 按这个写 https://juejin.cn/post/7093035836689612836
var JwtKey = []byte("yijiansanlian")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email  string `json:"email"`
	PassWord string `json:"pass_word"`
	OperationType uint    `json:"operation_type"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, username string, authority int) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	claims := Claims{
		ID:        id,
		UserName:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "FanOne-mall",
		},
	}

	// token 加密 签发
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtKey)
	return token, err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// EmailClaims 生成token
func GenerateEmailToken(userid uint, password string, email string, operationtype uint) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	claims := EmailClaims{
		UserID: userid,
		PassWord: password,
		Email: email,
		OperationType: operationtype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "email",
		},
	}

	// token 加密 签发
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtKey)
	return token, err
}

// EmailClaims 解码验证Token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tokenClaims != nil {
		if Emailclaims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return Emailclaims, nil
		}
	}

	return nil, err
}