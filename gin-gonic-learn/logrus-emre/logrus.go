package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.WithField("Info", "CreateFile").Info("starting file c")
	logrus.WithField("Debug", "CreateFile").Debug("debug file")
	f, err := os.Create("logrus.log")

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "CreateFile",
			"error":  err,
		}).Error(err.Error())
	}

	logrus.WithField("Info", "CreateFile").Debug("end file")
	router := gin.New()

	/*
		log.WithFields(log.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
		}).Info("A walrus appears")

	*/

	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	multi := io.MultiWriter(f, os.Stdout)

	logrus.SetOutput(multi)

	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
	})

	//logrus.SetOutput(os.Stdout)

	logrus.Traceln("Trace")
	logrus.Debugln("Debug")
	logrus.Warnln("Warn")
	logrus.Errorln("Error")

	router.Run()

}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{

		"data": "this is get data method",
	})

}
