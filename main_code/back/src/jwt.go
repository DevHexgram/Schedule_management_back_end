package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Authority int `json:"authority"`
	jwt.StandardClaims
}

var jwtSecret = []byte("lmx-Hexagram_and_legends-killer")

func GenerateToken(username string,authority int,id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * time.Hour * 24)
	strInt := strconv.Itoa(int(id))

	claims := Claims{
		Username: username,
		Authority:authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Id:strInt,
			Issuer:    "Sch_man",
		},
	}
	fmt.Println(claims)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	//fmt.Println(tokenClaims)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims,nil
		}
	}

	return nil,err
}
