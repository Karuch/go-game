package server

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/yalp/jsonpath"
  "encoding/json"
)

func Server() {
  r := gin.Default()
  r.POST("/ping", func(c *gin.Context) {
    // Read the request body
    requestBody, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Print the request body
    allAuthors, err := jsonpath.Prepare("$.guess")
    var bookstore interface{}
    err = json.Unmarshal(requestBody, &bookstore)
    authors, err := allAuthors(bookstore)
    fmt.Println(authors)
    // Respond with a message
    c.JSON(http.StatusOK, gin.H{"message": "Request body received successfully"})
  }) 
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}