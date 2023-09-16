package rest

import (
	userHandler "github.com/edunota/backend/internal/app/rest/handlers/user"
	"github.com/gin-gonic/gin"
)

func New( /* inject any dependencies from here*/ ) *gin.Engine {
	engine := gin.Default()

	router := engine.Group("/api/rest/v1")
	userHandler.UserHandler(router.Group("/users"))

	return engine
}
