package users

import (
	"main/router"
	"main/users/application/services"
	"main/users/domain/factories"
	"main/users/infrastructures/persistence/repositories"
	"main/users/presenters/http/controllers"
	"main/uuid"
)

func GetService() *services.UsersService {
	return &services.UsersService{
		UsersFactory: &factories.UsersFactory{
			UuidService: uuid.GetService(),
		},
		UsersRepository: &repositories.FakeUsersRepository{},
	}
}

func Init(mainRouter *router.MainRouter) {
	usersController := &controllers.UsersController{
		UsersService: GetService(),
	}

	mainRouter.Get("/users", usersController.GetAll)
	mainRouter.Get("/users/{id}", usersController.GetById)
	mainRouter.Post("/users", usersController.Create)
	mainRouter.Put("/users/{id}", usersController.Update)
	mainRouter.Delete("/users/{id}", usersController.Delete)
}