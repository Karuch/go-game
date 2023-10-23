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
  
  r.POST("/guess", func(c *gin.Context) {	
    requestBody, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    allAuthors, err := jsonpath.Prepare("$.guess")
    var bookstore interface{}
    err = json.Unmarshal(requestBody, &bookstore)
    authors, err := allAuthors(bookstore)
    if client_input, ok := authors.(string); ok { //str is the input from client.go
      doNothurtHM, listState, info := HangmanHandler(inVisable_array, visable_array, word, client_input) //DoNotHurtHM false = hit him
      c.String(http.StatusOK, DrawHM(doNothurtHM, listState, info)) //will return to clien the result of hangmanhandler
    } else { // Handle the case where authors is not a string.
      fmt.Println("Conversion to string failed")
    }
    
    
  }) 
  
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}