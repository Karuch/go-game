package main

import (
	//"bufio"
	"client/functions"
	//"fmt"
	//"net"
	//"os"
	//"strings"
)

func main() {

    //client.WriteToFile("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBY2Nlc3NUb2tlbjMiLCJleHAiOjE2MzAwMDAwMDB9.RS5MZ1l4Ym9Qd2dKZXBTMW5PVHJQZz10", ".\\file.txt")
    //fmt.Println(client.ReadFile(".\\AccessToken.txt"))
    //fmt.Println(client.ReadFile(".\\RefreshToken.txt"))
    client.HttpGetHandlerFIX("http://localhost:8080/health", client.ReadFile(".\\AccessToken.txt"))






   
}
