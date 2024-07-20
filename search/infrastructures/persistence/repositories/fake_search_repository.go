package repositories

import (
	"encoding/json"
	"fmt"
)

type FakeSearchRepository struct{}

type Document struct{
	Id string
	Body struct{}
}
var indices map[string][]*Document = map[string][]*Document{
	"users": {},
}

func ResetFakeSearchRepository() {
	indices = map[string][]*Document{
		"users": {},
	}
}

func (fsr *FakeSearchRepository) Create(index string, documentId string, body []byte) error {
	_, ok := indices[index]
	if !ok {
		return fmt.Errorf("index %s does not exist", index)
	}

	var decodedBody struct{}
	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		return fmt.Errorf("could not unmarshall body")
	}

	document := &Document{
		Id: documentId,
		Body: decodedBody,
	}
	indices[index] = append(indices[index], document)

	return nil
}

func (fsr *FakeSearchRepository) Update(index string, documentId string, body []byte) error {
	_, ok := indices[index]
	if !ok {
		return fmt.Errorf("index %s does not exist", index)
	}

	for _, document := range indices[index] {
		if document.Id == documentId {
			var decodedBody struct{}
			err := json.Unmarshal(body, &decodedBody)
			if err != nil {
				return fmt.Errorf("could not unmarshall body")
			}

			document.Body = decodedBody
		}
	}

	return nil
}

func (fsr *FakeSearchRepository) Delete(index string, documentId string) error {
	_, ok := indices[index]
	if !ok {
		return fmt.Errorf("index %s does not exist", index)
	}

	for documentIndex, document := range indices[index] {
		if document.Id == documentId {
			indices[index] = append(indices[index][:documentIndex], indices[index][documentIndex + 1:]...)
			break
		}
	}

	return nil
}