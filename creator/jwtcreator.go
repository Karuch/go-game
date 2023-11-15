package creator

import (
	"fmt"
	"log"
	"test/service"
	"time"
    "strconv"
	"github.com/golang-jwt/jwt"
)

func Creator(session_id int) {
    // user login validation should occur here
    // Assuming session_id is an integer
    var session_id_str = strconv.Itoa(session_id)

    // Convert session_id to string if it's not a string

    // Assign session_id to the Id field in UserClaims

    userClaims := service.UserClaims{
        Id:    session_id_str,
        //First: "Leeroy",
        //Last:  "Jenkins",
        StandardClaims: jwt.StandardClaims{
            IssuedAt:  time.Now().Unix(),
            ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
        },
    }

    signedAccessToken, err := service.NewAccessToken(userClaims)
    if err != nil {
        log.Fatal("error creating access token")
    }
    fmt.Println("ACCESS:  "+signedAccessToken)
    refreshClaims := jwt.StandardClaims{
        IssuedAt:  time.Now().Unix(),
        ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
    }

    signedRefreshToken, err := service.NewRefreshToken(refreshClaims)
    if err != nil {
        log.Fatal("error creating refresh token")
    }
    fmt.Println("REFRESH:  "+signedRefreshToken)
    // do something with access, and refresh tokens.
    // i.e. return them in an HTTP response for a successful login
}
