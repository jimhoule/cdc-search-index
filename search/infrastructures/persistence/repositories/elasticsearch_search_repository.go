package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"main/searchclient"
)

type ElasticsearchSearchRepository struct{
	SearchClient *searchclient.SearchClient
}

func (esr *ElasticsearchSearchRepository) Search(index string, queryType string, key string, value string) {
	// Creates request body
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			queryType: map[string]interface{}{
				key: value,
			},
		},
	}

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// Executes request
	response, err := esr.SearchClient.Client.Search(
		esr.SearchClient.Client.Search.WithIndex(index),
		esr.SearchClient.Client.Search.WithBody(&buffer),
	)
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer response.Body.Close()

	// Gets response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("error: ", err)
	}

	//
	for _, hit := range responseBody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		document := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Println("document: ", document)
	}
}

func (esr *ElasticsearchSearchRepository) GetById(index string, documentId string) (any, error) {
	// Creates request
	request := searchclient.GetRequest{
		Index: index, 
		DocumentID: documentId,
	}

	// Executes request
	response, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Gets response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	fmt.Println("document: ", responseBody["_source"])

	return "", nil
}

func (esr *ElasticsearchSearchRepository) Create(index string, documentId string, body []byte) error {
	// Creates request
	request := searchclient.CreateRequest{
		Index:      index,
		DocumentID: documentId,
		Body:       bytes.NewReader(body),
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return err
	}

	return nil
}

func (esr *ElasticsearchSearchRepository) Update(index string, documentId string, body []byte) error {
	// Creates request
	request := searchclient.UpdateRequest{
		Index:      index,
		DocumentID: documentId,
		Body:       bytes.NewReader(body),
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return err
	}

	return nil
}

func (esr *ElasticsearchSearchRepository) Delete(index string, documentId string) error {
	// Creates request
	request := searchclient.DeleteRequest{
		Index:      index,
		DocumentID: documentId,
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return err
	}

	return nil
}