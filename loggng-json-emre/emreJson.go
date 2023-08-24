package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	f, _ := os.Create("logging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router.GET("/getData", getData)

	router.Run()

}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "this is getting data",
	})

}
