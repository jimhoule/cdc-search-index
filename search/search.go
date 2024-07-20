package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Client = *elasticsearch.Client

type SearchHandler struct{
	Client Client
}

func NewSearchHandler() *SearchHandler {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println("error: ", err)
	}

	return &SearchHandler{
		Client: esClient,
	}
}

func (sh *SearchHandler) CreateIndex() {
	request := esapi.IndexRequest{
		Index: "",
		DocumentID: "",
		Body: strings.NewReader(""),
	}

	response, err := request.Do(context.Background(), sh.Client)
	if err != nil {
		fmt.Println("error: ", err)
	}
	dresponse.Body.Close()
}

func (sh *SearchHandler) GetDocument(index string, id string) {
	request := esapi.GetRequest{
		Index: index, 
		DocumentID: id,
	}

	response, err := request.Do(context.Background(), sh.Client)
	if err != nil {
		fmt.Println("error: ", err)
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	fmt.Println("document: ", result["_source"])
}

func (sh *SearchHandler) Search(index string, queryType string, key string, value string) {
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			queryType: map[string]interface{}{
				key: value,
			},
		},
	}

	json.NewEncoder(&buffer).Encode(query)

	response, _ := sh.Client.Search(
		sh.Client.Search.WithIndex(index),
		sh.Client.Search.WithBody(&buffer),
	)

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		document := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Println("document: ", document)
	}
}