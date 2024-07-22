package users

import (
	"fmt"
	"main/database"
	"main/queue"
	"main/router"
	"main/users/application/services"
	"main/users/domain/factories"
	"main/users/infrastructures/persistence/repositories"
	"main/users/presenters/http/controllers"
	"main/uuid"
	"os"
)

func GetService(db *database.Db) *services.UsersService {
	return &services.UsersService{
		UsersFactory: &factories.UsersFactory{
			UuidService: uuid.GetService(),
		},
		UsersRepository: &repositories.ArangodbUsersRepository{
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *database.Db) {
	queueProducerHandler, err := queue.NewProducerHandler(
		[]string{
			fmt.Sprintf("%s:%s", os.Getenv("QUEUE_URL"), os.Getenv("QUEUE_PORT")),
		},
	)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	usersController := &controllers.UsersController{
		UsersService:         GetService(db),
		QueueProducerHandler: &queueProducerHandler,
	}

	mainRouter.Get("/users", usersController.GetAll)
	mainRouter.Get("/users/{id}", usersController.GetById)
	mainRouter.Post("/users", usersController.Create)
	mainRouter.Put("/users/{id}", usersController.Update)
	mainRouter.Delete("/users/{id}", usersController.Delete)
}
