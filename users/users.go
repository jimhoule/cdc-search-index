package users

import (
	"fmt"
	"main/queue"
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
	queueProducerHandler, err := queue.NewProducerHandler([]string{"localhost:9092"})
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	usersController := &controllers.UsersController{
		UsersService: GetService(),
		QueueProducerHandler: &queueProducerHandler,
	}

	mainRouter.Get("/users", usersController.GetAll)
	mainRouter.Get("/users/{id}", usersController.GetById)
	mainRouter.Post("/users", usersController.Create)
	mainRouter.Put("/users/{id}", usersController.Update)
	mainRouter.Delete("/users/{id}", usersController.Delete)
}