package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/getData", getData)

	router.Run()

}

func getData(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data ": "Hi i am gin framework ",
	})
}
