package repositories

import (
	"encoding/json"
	"fmt"
	"main/search/domain/indices"
)

type FakeSearchRepository[T any] struct{}

type Document struct {
	Id   string
	Body any
}

var documents map[string][]*Document = map[string][]*Document{
	indices.UsersIndex: {},
}

func ResetFakeSearchRepository() {
	documents = map[string][]*Document{
		indices.UsersIndex: {},
	}
}

func (fsr *FakeSearchRepository[T]) GetAllByIndex(index string) ([]*T, error) {
	// Checks if index exists
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	views := []*T{}
	for _, document := range documents[index] {
		view := document.Body.(T)
		views = append(views, &view)
	}

	return views, nil
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
			view := document.Body.(T)
			return &view, nil
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

	var view T
	err := json.Unmarshal(body, &view)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall body")
	}

	document := &Document{
		Id:   documentId,
		Body: view,
	}
	documents[index] = append(documents[index], document)

	return &view, nil
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
			var view T
			err := json.Unmarshal(body, &view)
			if err != nil {
				return nil, fmt.Errorf("could not unmarshall body")
			}

			document.Body = view
			return &view, nil
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
