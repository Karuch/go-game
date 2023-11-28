package main

import (
	//"main/internal/authserver"
	//"main/internal/authserver/jwthandling/creator"
	"fmt"
	
	//main/internal/authserver/jwthandling/service"
	"os"
	"main/internal/authserver/jwthandling/validator"
)

func main(){
    os.Setenv("TOKEN_SECRET", "foo")
    
    //token, err := service.ParseAccessToken("j.j")
    //if err != nil {
    //    return 
    //}
    //fmt.Println("here", token)

    //fmt.Println(test.Name)
    //test := service.ParseAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.6Vce4TLYaUEExYEVgFtUFSx5_OkRsMlzmTKiRAlcwHA")
    //fmt.Println(test.Name)
    valid := validator.Validator_refreshtoken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJleGFtcGxlLmNvbSIsInN1YiI6InVzZXIiLCJpYXQiOjE3MDExNTk1MzgsImV4cCI6MTcwMTI0NTkzOH0.BfnenNCQp_zHqpnE4M1RXzBAAQkv-d2NROO8RT-IDiQ")
    fmt.Println(valid)
    //creator.Creator(2)
    //authserver.AuthServer()
    //server.Server()
}