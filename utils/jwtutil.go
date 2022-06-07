package utils

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"log"
	"time"
)

type CustomClaims struct {
	model.Person
	jwt.StandardClaims
}

var MySecret = []byte("Yezzi")

// 创建 Token
func GenToken(user model.Person) (string, error) {

	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 15)),
			Issuer:    "Yezzi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	log.Panicln(base.InvalidToken, err)
	return nil, errors.New("invalid token")
}

// 刷新 Token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return GenToken(claims.Person)
	}
	return "", errors.New("Cloudn't handle this token")
}
