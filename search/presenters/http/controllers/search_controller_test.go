package controllers

import (
	"encoding/json"
	"fmt"
	"main/router/mock"
	"main/search/application/payloads"
	"main/search/application/services"
	"main/search/domain/indices"
	"main/search/domain/views"
	"main/search/infrastructures/persistence/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*SearchController[views.UserView], func(), func() (*views.UserView, error)) {
	searchController := &SearchController[views.UserView]{
		SearchService: &services.SearchService[views.UserView]{
			SearchRepository: &repositories.FakeSearchRepository[views.UserView]{},
		},
	}

	userView := &views.UserView{
		Id:        "dummyDocumentId",
		Firstname: "dummy firstname",
		Lastname:  "dummy lastname",
	}
	body, _ := json.Marshal(userView)

	create := func() (*views.UserView, error) {
		return searchController.SearchService.Create(&payloads.CreatePayload{
			Index:      indices.UsersIndex,
			DocumentId: userView.Id,
			Body:       body,
		})
	}

	return searchController, repositories.ResetFakeSearchRepository, create
}

func TestGetAllByIndexSearchController(t *testing.T) {
	searchController, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search/%s", indices.UsersIndex), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"index": indices.UsersIndex,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(searchController.GetAllByIndex)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var userViews []*views.UserView
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &userViews)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *userViews[0] != *newUserView {
		t.Errorf("Expected first element of UserViews slice to equal NewUserView but got %v", *userViews[0])
	}
}

func TestGetByDocumentIdSearchController(t *testing.T) {
	searchController, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search/%s/%s", indices.UsersIndex, newUserView.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"index": indices.UsersIndex,
		"documentId": newUserView.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(searchController.GetByDocumentId)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var userView *views.UserView
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &userView)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *userView != *newUserView {
		t.Errorf("Expected UserView to equal NewUserView but got %v", *userView)
	}
}