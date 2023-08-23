package middleware

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {

	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "token onaylanmadı",
		})
		return
	}
	c.Next()

}
