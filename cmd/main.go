package main

import (
	"github.com/edunota/backend/internal/app/rest"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", rest.GetHelloWorld)
	return r
}

func main() {

	router := setupRouter()

	router.Run()

}
