package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/iSatVeerSingh/gossip/utils"
)

type Claims struct {
	UserInfo interface{}
	jwt.RegisteredClaims
}

func GenerateToken(userInfo interface{}) string {
	secretKey := utils.GetEnv("SESSION_SECRET")

	claims := &Claims{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatal(err)
	}
	return tokenStr
}

func ValidateToken(tokenStr string) (interface{}, bool) {

	secretKey := utils.GetEnv("SESSION_SECRET")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", false
	}

	if !token.Valid {
		return "", false
	}

	return token.Claims.(*Claims).UserInfo, true
}
