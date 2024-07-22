package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"main/searchclient"
)

type ElasticsearchSearchRepository[T any] struct {
	SearchClient *searchclient.SearchClient
}

func (esr *ElasticsearchSearchRepository[T]) GetDocuments(index string, queryType string, key string, value string) ([]*T, error) {
	// Creates request body
	query := map[string]interface{}{
		"query": map[string]interface{}{
			queryType: map[string]interface{}{
				key: value,
			},
		},
	}

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	// Executes request
	response, err := esr.SearchClient.Client.Search(
		esr.SearchClient.Client.Search.WithIndex(index),
		esr.SearchClient.Client.Search.WithBody(&buffer),
	)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	// Gets response body
	var responseBody map[string]any
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	// Gets slices of *T to return
	bodies := []*T{}
	for _, hit := range responseBody["hits"].(map[string]any)["hits"].([]any) {
		//body := hit.(map[string]any)["_source"].(*T)
		var body *T
		transcode(hit.(map[string]any)["_source"], body)

		bodies = append(bodies, body)
	}

	return bodies, nil
}

func (esr *ElasticsearchSearchRepository[T]) GetByDocumentId(index string, documentId string) (*T, error) {
	// Creates request
	request := searchclient.GetRequest{
		Index:      index,
		DocumentID: documentId,
	}

	// Executes request
	response, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Gets response body
	var responseBody map[string]any
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	var view T
	transcode(responseBody["_source"], &view)

	return &view, nil
}

func (esr *ElasticsearchSearchRepository[T]) Create(index string, documentId string, body []byte) (*T, error) {
	// Creates request
	request := searchclient.CreateRequest{
		Index:      index,
		DocumentID: documentId,
		Body:       bytes.NewReader(body),
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return nil, err
	}

	var view T
	json.Unmarshal(body, &view)

	return &view, nil
}

func (esr *ElasticsearchSearchRepository[T]) Update(index string, documentId string, body []byte) (*T, error) {
	// Creates request
	request := searchclient.UpdateRequest{
		Index:      index,
		DocumentID: documentId,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{ "doc": %s }`, body))),
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return nil, err
	}

	var view T
	json.Unmarshal(body, &view)

	return &view, nil
}

func (esr *ElasticsearchSearchRepository[T]) Delete(index string, documentId string) (string, error) {
	// Creates request
	request := searchclient.DeleteRequest{
		Index:      index,
		DocumentID: documentId,
	}

	// Executes request
	_, err := request.Do(context.Background(), esr.SearchClient.Client)
	if err != nil {
		return "", err
	}

	return documentId, nil
}

func transcode(in any, out any) {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(in)
	json.NewDecoder(&buffer).Decode(out)
}
