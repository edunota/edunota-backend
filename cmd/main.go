package main

import (
	"github.com/edunota/backend/internal/app/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	router := InitRouter()
	router.Run(":8081")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", rest.GetHelloWorld)
	return r
}
