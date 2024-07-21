package search

import (
	"fmt"
	"main/queue"
	"main/router"
	"main/search/application/services"
	"main/search/domain/views"
	"main/search/infrastructures/persistence/repositories"
	"main/search/presenters/consumer/handlers"
	"main/search/presenters/http/controllers"
	"main/searchclient"
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

	mainRouter.Get("/search/{index}/{documentId}", searchController.GetByDocumentId)

	// Creates queue consumer handler
	usersHandler := &handlers.UsersHandler[views.UserView]{
		SearchService: searchService,
	}

	usersConsumerGroupHandler, err := queue.NewConsumerGroupHandler([]string{"localhost:9092"}, "users_consumer_group")
	if err != nil {
		fmt.Println("error: ", err)
	}

	usersConsumerGroupHandler.Handlers = map[string]queue.Handler{
		"user.created": usersHandler.Create,
		"user.updated": usersHandler.Update,
		"user.deleted": usersHandler.Delete,
	}

	go func() {
		err := usersConsumerGroupHandler.Listen([]string{"user.created", "user.updated", "user.deleted"})
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()
}
