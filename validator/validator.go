package validator

import (
    "test/service"
    "log"
    "fmt"
    "github.com/golang-jwt/jwt"
    "time"
)

//var request_token string =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDAyMDkyMjIsImlhdCI6MTcwMDAzNjQyMn0.lwwUmZxzGSBDfybffSZVRmaNgEKt4I5nAwzCgRrNaOg" 

func Validator_refreshtoken(request_token string) {
    refreshClaims := service.ParseRefreshToken(request_token)

    var err error

    if validErr := refreshClaims.Valid(); validErr == nil {
        // Refresh token is valid
        fmt.Println("REFRESH TOKEN IS VALIDDDDDDDDDDDDDDDDDDDDDDDD")
    } else {
        // Refresh token is not valid
        fmt.Println("REFRESH TOKEN IS NOT VALIDDDDDDDDDDDDDDDD")
        request_token, err = service.NewRefreshToken(*refreshClaims)
        if err != nil {
            log.Fatal("error creating refresh token")
        }
    }
    

    // Additional logic for access token can be included here if needed
}


func Validator_accesstoken(access_token string) {
  accessClaims := service.ParseAccessToken(access_token)

  var err error

  if validErr := accessClaims.Valid(); validErr == nil {
      // Refresh token is valid
      fmt.Println("ACCESS TOKEN IS VALIDDDDDDDDDDDDDDDDDDDDDDDD")
  } else {
      // Refresh token is not valid
      fmt.Println("ACCESS TOKEN NOT VALIDDDDDDDDDDDDDDDD")
      access_token, err = service.NewAccessToken(*accessClaims)
      if err != nil {
          log.Fatal("error creating access token")
      }
  }
}

func Translator(token_input string){  
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(token_input, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("TOKEN_SECRET"), nil
    })
    // ... error handling
    if err != nil {
        fmt.Println(err)
    }
    // do something with decoded claims
    for key, val := range claims {
        fmt.Printf("Key: %v, value: %v\n", key, val)
    }
    fmt.Println(token)
}

func Time(){
        // JWT expiration timestamp
        expTimestamp := int64(1700046484)

        // Convert Unix timestamp to time.Time
        expTime := time.Unix(expTimestamp, 0)
    
        // Print the date in a human-readable format
        fmt.Println("Expiration date:", expTime.Format(time.DateTime))
}