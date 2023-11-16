package server

import ( 
	"github.com/gin-gonic/gin"
	"net/http"
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