package userModule

type IUserModule interface {
}

type userModule struct{}

func New() IUserModule {
	return &userModule{}
}
