package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.GET("/getData", getData)

	router.GET("/getQuerySettings", getQuerySettings)

	router.GET("/getUrlData/:name/:age", getUrlData)

	router.POST("/postData", postData)

	router.Run()

}

func postData(c *gin.Context) {

	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	c.JSON(200, gin.H{
		"deneme": "1",
		"body":   string(value),
	})

}

// http://localhost:8080/getUrlData/emre/80

func getUrlData(c *gin.Context) {

	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"data ": "burası url yeri",
		"name":  name,
		"age":   age,
	})

}

// http://localhost:8080/getQuerySettings?age=30&name=emre
func getQuerySettings(c *gin.Context) {

	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{

		"data ": "burası query yeri",
		"name ": name,
		"age":   age,
	})

}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{
		"data ": "Hi i am gin framework emre",
	})
}
