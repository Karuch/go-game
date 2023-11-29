package main

import (
	//"bufio"
	"client/functions"
	"fmt"
	"strings"
	//"encoding/json"
	//"net"
	//"os"
	//"strings"
)

func main() {

    //client.WriteToFile("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBY2Nlc3NUb2tlbjMiLCJleHAiOjE2MzAwMDAwMDB9.RS5MZ1l4Ym9Qd2dKZXBTMW5PVHJQZz10", ".\\file.txt")
    //fmt.Println(client.ReadFile(".\\AccessToken.txt"))
    //fmt.Println(client.ReadFile(".\\RefreshToken.txt"))
    answer := client.HttpGetHandlerFIX("http://localhost:8080/health", client.ReadFile(".\\AccessToken.txt"))
	fmt.Println("printing the answer client side:" + answer)
	if strings.Contains(answer, "needAccessToken") {
		client.HttpGetHandlerFIX("http://localhost:8080/health", "accessTokenPlease")
	} else if strings.Contains(answer, "needRefreshToken") {
		client.HttpGetHandlerFIX("http://localhost:8080/health", "refreshTokenPlease")
	} else if strings.Contains(answer, "GoodAcc") { //instead of GoodAccessToken problem should fix Karuch
		fmt.Println("answer:", answer, "You are good.")
	} else { 
		fmt.Println("else should help debug" + answer) //should never accoured if yes unknown behavior of software
	}




   
}
