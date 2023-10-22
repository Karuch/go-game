package server

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/yalp/jsonpath"
  "encoding/json"
  "strings"
)

var AlreadyFoundLettersSlice []string

func Server() {
  if AlreadyFoundLettersSlice == nil {
    fmt.Println("is nilllllll")
    AlreadyFoundLettersSlice = []string{}
  }
	//initiallize blank and word list
	var word string = "He llo"
	visable_array := strings.Split(word, "")
	inVisable_array := make([]string, len(visable_array))
	copy(inVisable_array[:], visable_array[:])
	for i := 0; i < len(inVisable_array); i++ {
    if inVisable_array[i] != " " {
		  inVisable_array[i] = "_"
    }
	}
	fmt.Println(inVisable_array, visable_array)

  //server
  r := gin.Default()
  
  r.POST("/ping", func(c *gin.Context) {
    AlreadyFoundLettersSlice = append(AlreadyFoundLettersSlice, "t")
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
    if str, ok := authors.(string); ok {
      // Conversion was successful
      // Now, str is a string containing the value.
      fmt.Println(str)
      
      _, _, three := HangmanHandler(inVisable_array, visable_array, word, str, AlreadyFoundLettersSlice)
      fmt.Println(three)
    } else {
      // Handle the case where authors is not a string.
      fmt.Println("Conversion to string failed")
    }
    // Respond with a message
    c.JSON(http.StatusOK, gin.H{"message": "Request body received successfully"}) //will return to clien the result of hangmanhandler
  }) 
  
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}