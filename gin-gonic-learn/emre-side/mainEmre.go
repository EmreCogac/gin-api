package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// router.GET("/getData", getData) // admine atadım

	// router.GET("/getQuerySettings", getQuerySettings) // client a atadım

	router.GET("/getUrlData/:name/:age", getUrlData)

	router.POST("/postData", postData)

	auth := gin.BasicAuth(gin.Accounts{
		"user":   "pass",
		"user2 ": "pass2",
		"user3 ": "pass3",
	})

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}

	client := router.Group("/client")
	{
		client.GET("/getQuerySettings", getQuerySettings)
	}

	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()

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
