package search

import (
	"fmt"
	"main/queue"
	"main/queue/topics"
	"main/router"
	"main/search/application/services"
	"main/search/domain/views"
	"main/search/infrastructures/persistence/repositories"
	"main/search/presenters/consumer/handlers"
	"main/search/presenters/http/controllers"
	"main/searchclient"
	"os"
)

func Init(mainRouter *router.MainRouter, searchClient *searchclient.SearchClient) {
	// Creates service
	searchService := &services.SearchService[views.UserView]{
		SearchRepository: &repositories.ElasticsearchSearchRepository[views.UserView]{
			SearchClient: searchClient,
		},
	}

	// Creates http routes
	searchController := &controllers.SearchController[views.UserView]{
		SearchService: searchService,
	}

	mainRouter.Get("/search/{index}", searchController.GetAllByIndex)
	mainRouter.Get("/search/{index}/{documentId}", searchController.GetByDocumentId)

	// Creates queue consumer handler
	usersHandler := &handlers.UsersHandler[views.UserView]{
		SearchService: searchService,
	}

	usersConsumerGroupHandler, err := queue.NewConsumerGroupHandler(
		[]string{
			fmt.Sprintf("%s:%s", os.Getenv("QUEUE_URL"), os.Getenv("QUEUE_PORT")),
		},
		"users_consumer_group_id",
	)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// Associates each handler to proper topic
	usersConsumerGroupHandler.Handlers = map[string]queue.Handler{
		topics.UserCreated: usersHandler.Create,
		topics.UserUpdated: usersHandler.Update,
		topics.UserDeleted: usersHandler.Delete,
	}

	go func() {
		err := usersConsumerGroupHandler.Listen(
			[]string{
				topics.UserCreated,
				topics.UserUpdated,
				topics.UserDeleted,
			},
		)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()
}
