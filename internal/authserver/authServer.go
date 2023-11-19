package authserver

import ( 
	"github.com/gin-gonic/gin"
	"net/http"
    "fmt"
)

type RequestData struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func AuthServer() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "looks good.",
		})
        fmt.Println("auth:", c.GetHeader("Authorization"))
        if c.GetHeader("Authorization") != "Bearer 1" {
            fmt.Println("ACCESS TOKEN IS VALID | WILL ALLOW SESSION ACCESS")
        } else if c.GetHeader("Authorization") != "test" {
            fmt.Println("REFRESH TOKEN IS VALID | WILL GENEREATE ACCESS TOKEN WITH REFRESH ID THEN ASK CLIENT FOR ACCESS")
        } else if "" == "" { //case client was not found any token in it's then ask server for new refresh
            fmt.Println("WILL GENERATE TOKEN")
        }
	})
    r.POST("/", func(c *gin.Context) {
        var data RequestData

        // Bind the JSON to the struct
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Process the data (just printing here)
        c.JSON(http.StatusOK, gin.H{"message": "Data received", "name": data.Name, "age": data.Age})
        
    })

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getAllHeaders(c *gin.Context) map[string]string {
    headers := make(map[string]string)
    for key, values := range c.Request.Header {
        headers[key] = values[0] // If you expect multiple values per key, adjust this accordingly
    }
    return headers
}