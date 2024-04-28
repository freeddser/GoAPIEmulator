package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetTimeOutByNumber(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		number = 10
	}
	fmt.Println("It will sleep " + strconv.Itoa(number) + "s")
	time.Sleep(time.Duration(number) * time.Second)
	fmt.Println("It sleep done:" + strconv.Itoa(number) + "s")
	c.JSONP(http.StatusOK, gin.H{
		"msg":      "success",
		"response": "It sleep done:" + strconv.Itoa(number) + "s",
	})
}
