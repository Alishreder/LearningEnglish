package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {


	router := gin.Default()
	router.GET("/", Default)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Cant listen and serve on 0.0.0.0:8080")
		return
	}
}

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "authorization.html", nil)
}
