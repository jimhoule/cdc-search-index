package repositories

import (
	"bytes"
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
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	for _, document := range documents[index] {
		if document.Id == documentId {
			// NOTE: This transcoding trick allows us to use a generic type
			var body T
			var buffer bytes.Buffer
			json.NewEncoder(&buffer).Encode(document.Body)
			json.NewDecoder(&buffer).Decode(&body)

			return &body, nil
		}
	}

	return nil, nil
}

func (fsr *FakeSearchRepository[T]) Create(index string, documentId string, body []byte) (*T, error) {
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	var decodedBody map[string]any
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall body")
	}

	document := &Document{
		Id:   documentId,
		Body: decodedBody,
	}
	documents[index] = append(documents[index], document)

	return nil, nil
}

func (fsr *FakeSearchRepository[T]) Update(index string, documentId string, body []byte) (*T, error) {
	_, ok := documents[index]
	if !ok {
		return nil, fmt.Errorf("index %s does not exist", index)
	}

	for _, document := range documents[index] {
		if document.Id == documentId {
			var decodedBody map[string]any
			err := json.Unmarshal(body, &decodedBody)
			if err != nil {
				return nil, fmt.Errorf("could not unmarshall body")
			}

			document.Body = decodedBody
			break
		}
	}

	return nil, nil
}

func (fsr *FakeSearchRepository[T]) Delete(index string, documentId string) (string, error) {
	_, ok := documents[index]
	if !ok {
		return "", fmt.Errorf("index %s does not exist", index)
	}

	for documentIndex, document := range documents[index] {
		if document.Id == documentId {
			documents[index] = append(documents[index][:documentIndex], documents[index][documentIndex+1:]...)
			break
		}
	}

	return documentId, nil
}
