package handlers

import (
	"Go-APIEmulator/flags"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHeader(c *gin.Context) {
	apikey := c.GetHeader("x-api-key")
	if apikey != flags.ApiKey {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"msg":      "failed",
			"response": "apikey error",
		})
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"msg":      "success",
			"response": "api works",
		})
	}

}
