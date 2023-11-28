package validator

import (
    "main/internal/authserver/jwthandling/service"
    //"log"
    "fmt"
    "github.com/golang-jwt/jwt"
    "time"
)

//var request_token string =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDAyMDkyMjIsImlhdCI6MTcwMDAzNjQyMn0.lwwUmZxzGSBDfybffSZVRmaNgEKt4I5nAwzCgRrNaOg" 

func Validator_refreshtoken(request_token string) bool {
    refreshClaims, err := service.ParseRefreshToken(request_token)
    if err != nil {
        return false
    }
    //var err error
    if validErr := refreshClaims.Valid(); validErr == nil { //just check if refreshclaims return error or not if not
        // Refresh token is valid
        fmt.Println("REFRESH TOKEN IS VALIDDDDDDDDDDDDDDDDDDDDDDDD")
        return true
    } else {
        // Refresh token is not valid
        fmt.Println("REFRESH TOKEN IS NOT VALIDDDDDDDDDDDDDDDD")
        //request_token, err = service.NewRefreshToken(*refreshClaims)
        //if err != nil {
        //    log.Fatal("error creating refresh token")
        //}
        return false
    }
    

    // Additional logic for access token can be included here if needed
}


func Validator_accesstoken(access_token string) bool { 
  fmt.Println("hereeee1")
  accessClaims, err := service.ParseAccessToken(access_token)
  if err != nil {
    return false
  }
  fmt.Println("hereeee2")
  if validErr := accessClaims.Valid(); validErr == nil { //just check if accessclaims returns error or not
      // Refresh token is valid
      fmt.Println("ACCESS TOKEN IS VALIDDDDDDDDDDDDDDDDDDDDDDDD")
      return true
  } else { 
      // Refresh token is not valid
      fmt.Println("ACCESS TOKEN NOT VALIDDDDDDDDDDDDDDDD")
      //access_token, err = service.NewAccessToken(*accessClaims)
      //if err != nil {
      //    log.Fatal("error creating access token")
      //}
      return false
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
        expTimestamp := int64(1.700047972e+09)

        // Convert Unix timestamp to time.Time
        expTime := time.Unix(expTimestamp, 0)
    
        // Print the date in a human-readable format
        fmt.Println("Expiration date:", expTime.Format(time.DateTime))
}