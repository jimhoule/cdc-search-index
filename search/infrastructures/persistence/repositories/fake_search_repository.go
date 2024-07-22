package repositories

import (
	"encoding/json"
	"fmt"
)

type FakeSearchRepository[T any] struct{}

type Document struct {
	Id   string
	Body any
}

var documents map[string][]*Document = map[string][]*Document{
	"users": {},
}

func ResetFakeSearchRepository() {
	documents = map[string][]*Document{
		"users": {},
	}
}

func (fsr *FakeSearchRepository[T]) GetByDocumentId(index string, documentId string) (*T, error) {
	// Checks if index exists
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	for _, document := range documents[index] {
		// Checks if document exists
		if document.Id == documentId {
			body := document.Body.(T)
			return &body, nil
		}
	}

	return nil, nil
}

func (fsr *FakeSearchRepository[T]) Create(index string, documentId string, body []byte) (*T, error) {
	// Checks if index exists
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	var decodedBody T
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall body")
	}

	document := &Document{
		Id:   documentId,
		Body: decodedBody,
	}
	documents[index] = append(documents[index], document)

	return &decodedBody, nil
}

func (fsr *FakeSearchRepository[T]) Update(index string, documentId string, body []byte) (*T, error) {
	// Checks if index exists
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	for _, document := range documents[index] {
		// Checks if document exists
		if document.Id == documentId {
			var decodedBody T
			err := json.Unmarshal(body, &decodedBody)
			if err != nil {
				return nil, fmt.Errorf("could not unmarshall body")
			}

			document.Body = decodedBody
			return &decodedBody, nil
		}
	}

	return nil, fmt.Errorf("document with id %s does not exist", documentId)
}

func (fsr *FakeSearchRepository[T]) Delete(index string, documentId string) (string, error) {
	// Checks if index exists
	_, ok := documents[index]
	if !ok {
		return "", fmt.Errorf("index %s does not exist", index)
	}

	for documentIndex, document := range documents[index] {
		// Checks if document exists
		if document.Id == documentId {
			documents[index] = append(documents[index][:documentIndex], documents[index][documentIndex+1:]...)
			return documentId, nil
		}
	}

	return documentId, fmt.Errorf("document with id %s does not exist", documentId)
}
