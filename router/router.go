package router

import (
	handlers "Go-APIEmulator/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// ping
	r.GET("/", handlers.Hello)
	r.GET("/ping", handlers.Ping)

	//// time
	r.GET("/timeout/:number", handlers.GetTimeOutByNumber)
	//// GetHeader
	r.GET("/header", handlers.GetHeader)

	return r
}
