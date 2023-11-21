package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGetHandler(url string) {
	// Sending the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//resp.Header.Add("Authorization", "Bearer "+"TOKEN")
	defer resp.Body.Close() // Don't forget to close the response body
	// Reading the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	// Printing the response body
	fmt.Println(string(body))
}

func HttpGetHandlerFIX(url string, token string) string {
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "error"
	}

	// Add the authorization header to the request
	req.Header.Add("Authorization", token)

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "error"
	}
	defer resp.Body.Close() // Don't forget to close the response body

	// Reading the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return "error"
	}

	// Printing the response body
	return string(body)
}

func HttpPostHandler() {
	type RequestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// URL of the server endpoint
	url := "http://localhost:8080/"

	// Data to be sent in the POST request
	requestData := RequestData{
		Name: "Jane Doe",
		Age:  28,
	}

	// Convert struct to JSON
	jsonValue, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}

	fmt.Println("Response:", string(body))
}


func WriteToFile(content string, path string) { //create file automatically if dosen't exist ride content entirly
	data := []byte(content)
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("File writing error", err)
		return
	}
	fmt.Println("Data successfully written to file")
}

func ReadFile(path string) string {
	// get content from the file using ioutil package results will be in form of []bytes {10 30 40}
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
		if strings.Contains(path, "Access") {
			WriteToFile(HttpGetHandlerFIX("http://localhost:8080/health", "Access Failed"), ".\\AccessToken.txt") //here I'll write the token from the server default is None
			return ReadFile(path)
		} else if strings.Contains(path, "Refresh") {
			WriteToFile(HttpGetHandlerFIX("http://localhost:8080/health", "Refresh Failed"), ".\\RefreshToken.txt") //here I'll write the token from the server default is None
			return ReadFile(path)
		}
	}
	// results are in form of bytes [10 40 60]
	return string(fileContent)
}
