package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "I tested my first program using gin framework",
        })
    })
	
	r.POST("/blog", func(c *gin.Context) {
		var blog Blog
		err := c.ShouldBindJSON(&blog)
		if err != nil {
			errResponse:= gin.H{"error": err.Error()}
			c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
			return
		}
		response:= fmt.Sprintf("Blog created with title as %s",  blog.Title)
		c.JSON(200, gin.H{
			"response": response,
		})
	})

	r.Run()
}

type Blog struct {
	Title  	string `json:"title"`
	Section string `json:"section"`
}