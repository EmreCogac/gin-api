package main

import (
	"newproject/gin-api/middleware-emre/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.Use(middleware.Auth)

	admin := router.Group("/admin", middleware.Auth)
	{
		admin.GET("/getData", getData)
		admin.GET("/getData1", getData1)
		admin.GET("/getData2", getData2)
	}

	router.GET("/getData", getData)
	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)

	router.Run()

}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{
		"deneme": "deneme",
	})

}

func getData1(c *gin.Context) {

	c.JSON(200, gin.H{
		"deneme1": "deneme1",
	})

}

func getData2(c *gin.Context) {

	c.JSON(200, gin.H{
		"deneme 2": "deneme 2",
	})

}
