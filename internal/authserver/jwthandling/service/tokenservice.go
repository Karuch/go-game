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
        fmt.Println("Error parsing access token:", err)
        return nil, err //enoutred error, error != nil
    }

    return parsedAccessToken.Claims.(*UserClaims), err //error == nil
}

func ParseRefreshToken(refreshToken string) (*jwt.StandardClaims, error) {
    parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("TOKEN_SECRET")), nil
    })

    if err != nil {
        fmt.Println("Error parsing access token:", err)
        return nil, err //enoutred error, error != nil
    }

    return parsedRefreshToken.Claims.(*jwt.StandardClaims), err
}
