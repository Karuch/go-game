package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"bufio"
	"os"
)

func main(){
	fmt.Println("dwd")

	scanner := bufio.NewScanner(os.Stdin)

	client := resty.New()

	for 0 < 1 {
		scanner.Scan()
		test := scanner.Text()
		resp, err := client.R().
		SetBody(map[string]interface{}{"guess": test}).
		SetResult("Success").    // or SetResult(AuthSuccess{}).
		SetError("Error").       // or SetError(AuthError{}).
		Post("http://127.0.0.1:8080/ping")

		fmt.Println("Error:", err)
		fmt.Println(resp)
	}

	
	// Explore response object
	//fmt.Println("Response Info:")
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	
}