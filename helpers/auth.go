package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/iSatVeerSingh/gossip/utils"
)

type Claims struct {
	UserInfo utils.AuthUser
	jwt.RegisteredClaims
}

func GenerateToken(userInfo utils.AuthUser) string {
	secretKey := utils.GetEnv("SESSION_SECRET")

	claims := &Claims{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(utils.EXP_TIME)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatal(err)
	}
	return tokenStr
}

func ValidateToken(tokenStr string) (utils.AuthUser, bool) {

	secretKey := utils.GetEnv("SESSION_SECRET")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return token.Claims.(*Claims).UserInfo, false
	}

	if !token.Valid {
		return token.Claims.(*Claims).UserInfo, false
	}

	return token.Claims.(*Claims).UserInfo, true
}
