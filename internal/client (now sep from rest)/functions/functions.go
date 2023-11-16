package client

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

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
	}
	// results are in form of bytes [10 40 60]
	return string(fileContent)
}

func HttpGetHandler(url string){
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



func HttpGetHandlerFIX(url string, token string) {
    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Add the authorization header to the request
    req.Header.Add("Authorization", "Bearer " + "TOKEN")

    // Send the request using http.Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
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







func HttpPostHandler(){
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