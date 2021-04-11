package myjwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type jwtClaims struct {
	User
	jwt.StandardClaims
}

type User struct {
	Account string `json:"account"`
	Age     int    `json:"age"`
}

var (
	jwtKey = "jwtKeytest"
)

// CreateJwt create jwt token
func CreateJwt(user User) (string, error) {
	c := jwtClaims{
		User: User{
			Account: user.Account,
			Age:     user.Age,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 86400,
			Issuer:    "sgs",
			IssuedAt:  time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		log.Println("[jwt]CreateJwt error,", err)
	}
	//log.Println(token)
	return token, err
}

// ParseJwt parse jwt token
func ParseJwt(token string) (User, error) {
	if token == "" {
		return User{}, errors.New("no token found")
	}
	var c jwtClaims
	_, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		// assert断言
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		log.Println("[jwt]ParseJwt error,", err)
		return User{}, err
	}
	//log.Println(c)
	return c.User, nil
}

// CreateJwt create jwt token with Map
func CreateJwt2(u User) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Unix() + 30,
		"iss":     "sgs",
		"nbf":     time.Now().Unix() + 5,
		"account": "",
	})
	log.Println("token:", t)

	token, err := t.SignedString(jwtKey)
	if err != nil {
		log.Println("[jwt]CreateJwt error,", err)
	}
	log.Println(token)
}

// ParseJwt parse jwt token
func ParseJwt2(token string) {
	var m jwt.MapClaims
	_, err := jwt.ParseWithClaims(token, &m, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("[jwt]ParseJwt error,", err)
	}
	log.Println(m)
}
