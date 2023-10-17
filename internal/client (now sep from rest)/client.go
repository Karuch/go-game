package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/yalp/jsonpath"
	"encoding/json"
	"bufio"
	"os"
)

func main(){
	fmt.Println("dwd")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	test := scanner.Text()
	fmt.Println(test)

	client := resty.New()

	resp, err := client.R().EnableTrace().Get("http://127.0.0.1:8080/ping")

	// Explore response object
	//fmt.Println("Response Info:")
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	
	fmt.Println("Error:", err)
	allAuthors, err := jsonpath.Prepare("$.message")
	
	var bookstore interface{}
	err = json.Unmarshal(resp.Body(), &bookstore)
	authors, err := allAuthors(bookstore)
	fmt.Println(authors)

}