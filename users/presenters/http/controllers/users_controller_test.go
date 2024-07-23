package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	queueMock "main/queue/mock"
	routerMock "main/router/mock"
	"main/users/application/payloads"
	"main/users/application/services"
	"main/users/domain/factories"
	"main/users/domain/models"
	"main/users/infrastructures/persistence/repositories"
	"main/users/presenters/http/dtos"
	"main/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*UsersController, func(), func() (*models.User, error)) {
	queueProducerHandler := queueMock.NewMockProducerHandler()

	usersController := &UsersController{
		UsersService: &services.UsersService{
			UsersFactory: &factories.UsersFactory{
				UuidService: uuid.GetService(),
			},
			UsersRepository: &repositories.FakeUsersRepository{},
		},
		QueueProducerHandler: &queueProducerHandler,
	}

	createUser := func() (*models.User, error) {
		return usersController.UsersService.Create(&payloads.CreateUserPayload{
			Firstname: "Dummy firstname",
			Lastname:  "Dummy lastname",
		})
	}

	return usersController, repositories.ResetFakeUsersRepository, createUser
}

func TestCreateUserController(t *testing.T) {
	usersController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateUserDto{
		Firstname: "Dummy firstname",
		Lastname:  "Dummy lastname",
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(usersController.Create)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
	}
}

func TestGetAllUsersController(t *testing.T) {
	usersController, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(usersController.GetAll)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var users []*models.User
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &users)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *users[0] != *newUser {
		t.Errorf("Expected first element of Users slice to equal NewUser but got %v", *users[0])
	}
}

func TestGetUserByIdController(t *testing.T) {
	usersController, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", newUser.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newUser.Id,
	}
	request = routerMock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(usersController.GetById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var user *models.User
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *user != *newUser {
		t.Errorf("Expected User to equal NewUser but got %v", *user)
	}
}

func TestUpdateUpdateController(t *testing.T) {
	usersController, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	// Creates request body
	updatedFirstname := "Updated dummy firstname"
	requestBody, err := json.Marshal(dtos.UpdateUserDto{
		Firstname: updatedFirstname,
		Lastname: newUser.Lastname,
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/users/%s", newUser.Id), bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newUser.Id,
	}
	request = routerMock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(usersController.Update)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
	}

	// Validates response body
	var user *models.User
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if user.Firstname != updatedFirstname {
		t.Errorf("Expected User firstname to equal updated firstname but got %s", user.Firstname)
	}

	// Validates new user
	if newUser.Firstname != updatedFirstname {
		t.Errorf("Expected NewUser firstname to equal updated firstname but got %s", newUser.Firstname)
		return
	}
}

func TestDeleteUserController(t *testing.T) {
	usersController, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	// Creates request
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", newUser.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newUser.Id,
	}
	request = routerMock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(usersController.Delete)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusNoContent {
		t.Errorf("Expected http.StatusNoContent but got %d", responseRecorder.Code)
	}
}