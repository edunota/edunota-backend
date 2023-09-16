package userHandler

type IUserHandler interface {
	CreateUser()
	GetUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
}

type userHandler struct {
	userService any // replace it with actual interface
}

func (h *userHandler) CreateUser()     {}
func (h *userHandler) GetUsers()       {}
func (h *userHandler) GetUserById()    {}
func (h *userHandler) UpdateUserById() {}
func (h *userHandler) DeleteUserById() {}

func UserHandler() IUserHandler {
	return &userHandler{}
}
