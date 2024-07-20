package search

import (
	"fmt"
	"main/queue"
	"main/search/application/services"
	"main/search/infrastructures/persistence/repositories"
	"main/search/presenters/consumer/handlers"
	"main/searchclient"
)

func Init(searchClient *searchclient.SearchClient) {
	usersHandler := &handlers.UsersHandler{
		SearchService: &services.SearchService{
			SearchRepository: &repositories.ElasticsearchSearchRepository{
				SearchClient: searchClient,
			},
		},
	}

	consumerGroupHandler, err := queue.NewConsumerGroupHandler([]string{"localhost:9092"}, "users")
	if err != nil {
		fmt.Println("error: ", err)
	}

	consumerGroupHandler.Handlers["user.created"] = usersHandler.Create
	consumerGroupHandler.Handlers["user.updated"] = usersHandler.Update
	consumerGroupHandler.Handlers["user.deleted"] = usersHandler.Delete

	err = consumerGroupHandler.Listen([]string{"user.created", "user.updated", "user.deleted"})
	if err != nil {
		fmt.Println("error: ", err)
	}
}