package service

import (
    "github.com/golang-jwt/jwt"
    "os"
    "fmt"
)   

type UserClaims struct {
    Name    string `json:"name"`
    First string `json:"first"`
    Last  string `json:"last"`
    jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
} 

func ParseAccessToken(accessToken string) (*UserClaims, error) {
    parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("TOKEN_SECRET")), nil
    })
    if err != nil {
        fmt.Println(err)
    }
    return parsedAccessToken.Claims.(*UserClaims), err
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
    parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("TOKEN_SECRET")), nil
    })
    if err != nil {
        fmt.Println(err)
    }
    return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
