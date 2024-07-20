package searchclient

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Client = *elasticsearch.Client
type GetRequest = esapi.GetRequest
type CreateRequest = esapi.IndexRequest
type UpdateRequest = esapi.UpdateRequest
type DeleteRequest = esapi.DeleteRequest

type SearchClient struct {
	Client Client
}

var searchClient *SearchClient

func Get() *SearchClient {
	if searchClient == nil {
		esClient, err := elasticsearch.NewDefaultClient()
		if err != nil {
			fmt.Println("error creating search client: ", err)
			os.Exit(1)
		}

		searchClient = &SearchClient{
			Client: esClient,
		}
	}

	return searchClient
}