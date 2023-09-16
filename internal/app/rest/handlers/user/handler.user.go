package userHandler

import "github.com/gin-gonic/gin"

type userHandler struct {
	userService any // replace it with actual interface
}

func (h *userHandler) CreateUser(ctx *gin.Context)     {}
func (h *userHandler) GetUsers(ctx *gin.Context)       {}
func (h *userHandler) GetUserById(ctx *gin.Context)    {}
func (h *userHandler) UpdateUserById(ctx *gin.Context) {}
func (h *userHandler) DeleteUserById(ctx *gin.Context) {}

func UserHandler(r *gin.RouterGroup /* inject any other dependencies from here */) {
	h := userHandler{}
	r.POST("/", h.CreateUser)
	r.GET("/", h.GetUsers)
	r.GET("/:id", h.GetUserById)
	r.PATCH("/:d", h.UpdateUserById)
	r.DELETE("/:id", h.DeleteUserById)
}
