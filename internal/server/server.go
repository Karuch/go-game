package server

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/yalp/jsonpath"
  "encoding/json"
  "strings"
)

//GLOBAL
func Server() {
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

  //server
  r := gin.Default()
  
  r.POST("/ping", func(c *gin.Context) {
    
    requestBody, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    allAuthors, err := jsonpath.Prepare("$.guess")
    var bookstore interface{}
    err = json.Unmarshal(requestBody, &bookstore)
    authors, err := allAuthors(bookstore)
    if str, ok := authors.(string); ok { //str is the input from client.go
      DoNothurtHM, listState, info := HangmanHandler(inVisable_array, visable_array, word, str) //DoNotHurtHM false = hit him
      fmt.Println(DoNothurtHM, listState, info)
      fmt.Println(DrawHM(DoNothurtHM, listState, info))
    } else { // Handle the case where authors is not a string.
      fmt.Println("Conversion to string failed")
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Request body received successfully"}) //will return to clien the result of hangmanhandler
  }) 
  
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}