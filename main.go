package main

import (
	//"main/internal/authserver"
	//"main/internal/authserver/jwthandling/creator"
	//"fmt"
    "os"
	//"main/internal/authserver/jwthandling/service"
	"main/internal/authserver/jwthandling/validator"
)

func main(){
    os.Setenv("TOKEN_SECRET", "foo")
    //test := service.ParseAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.he0ErCNloe4J7Id0Ry2SEDg09lKkZkfsRiGsdX_vgEg")
    //fmt.Println(test.Name)
    //test := service.ParseAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.6Vce4TLYaUEExYEVgFtUFSx5_OkRsMlzmTKiRAlcwHA")
    //fmt.Println(test.Name)
    validator.Validator_accesstoken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Bm8tu4m18oA96xwhBL8AV_4hRpIU6OrK5UaOmGqBEsk")
    //creator.Creator(2)
    //authserver.AuthServer()
    //server.Server()
}