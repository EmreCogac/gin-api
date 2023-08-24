package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	router := gin.New()
	/*
		log.WithFields(log.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
		}).Info("A walrus appears")

	*/

	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stdout)
	/*
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Traceln("Trace")
		logrus.Debugln("Debug")
		logrus.Warnln("Warn")
		logrus.Errorln("Error")
		logrus.Fatalln("Fatal")
		router.GET("/getData", getData) */

	router.Run()

}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{

		"data": "this is get data method",
	})

}
