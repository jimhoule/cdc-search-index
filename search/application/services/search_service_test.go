package services

import (
	"encoding/json"
	"main/search/application/payloads"
	"main/search/domain/indices"
	"main/search/domain/views"
	"main/search/infrastructures/persistence/repositories"
	"testing"
)

func getTestContext() (*SearchService[views.UserView], func(), func() (*views.UserView, error)) {
	searchService := &SearchService[views.UserView]{
		SearchRepository: &repositories.FakeSearchRepository[views.UserView]{},
	}

	userView := &views.UserView{
		Id:        "dummyDocumentId",
		Firstname: "dummy firstname",
		Lastname:  "dummy lastname",
	}
	body, _ := json.Marshal(userView)

	create := func() (*views.UserView, error) {
		return searchService.Create(&payloads.CreatePayload{
			Index:      indices.UsersIndex,
			DocumentId: userView.Id,
			Body:       body,
		})
	}

	return searchService, repositories.ResetFakeSearchRepository, create
}

func TestCreateSearchService(t *testing.T) {
	_, reset, create := getTestContext()
	defer reset()

	_, err := create()
	if err != nil {
		t.Errorf("Expected to create a UserView but got %v", err)
		return
	}
}

func TestGetAllByIndexSearchService(t *testing.T) {
	searchService, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	userViews, err := searchService.GetAllByIndex(&payloads.GetByAllByIndexPayload{
		Index: indices.UsersIndex,
	})
	if err != nil {
		t.Errorf("Expected to get a UserViews by index but got %v", err)
		return
	}

	if len(userViews) != 1 {
		t.Errorf("Expected slice of UserViews with a length of 1 but got %d", len(userViews))
		return
	}

	if *userViews[0] != *newUserView {
		t.Errorf("Expected first UserView of slice to be equal to NewUserView but got %v", userViews[0])
		return
	}
}

func TestGetByDocumentIdSearchService(t *testing.T) {
	searchService, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	userView, err := searchService.GetByDocumentId(&payloads.GetByDocumentIdPayload{
		Index:      indices.UsersIndex,
		DocumentId: newUserView.Id,
	})
	if err != nil {
		t.Errorf("Expected to get a UserView by document id but got %v", err)
		return
	}

	if *newUserView != *userView {
		t.Errorf("Expected UserView to be equal to NewUserView but got %v", userView)
		return
	}
}

func TestUpdateSearchService(t *testing.T) {
	searchService, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	updatedUserView := &views.UserView{
		Id:        newUserView.Id,
		Firstname: "Updated dummy firstname",
		Lastname:  newUserView.Lastname,
	}
	body, _ := json.Marshal(updatedUserView)

	userView, err := searchService.Update(&payloads.UpdatePayload{
		Index:      indices.UsersIndex,
		DocumentId: newUserView.Id,
		Body:       body,
	})
	if err != nil {
		t.Errorf("Expected UserView but got %v", err)
		return
	}

	if userView.Firstname != updatedUserView.Firstname {
		t.Errorf("Expected UserView firstname to equal UpdatedUserView firstname but got %s", userView.Firstname)
	}
}

func TestDeleteSearchService(t *testing.T) {
	searchService, reset, create := getTestContext()
	defer reset()

	newUserView, _ := create()

	userViewId, err := searchService.Delete(&payloads.DeletePayload{
		Index:      indices.UsersIndex,
		DocumentId: newUserView.Id,
	})
	if err != nil {
		t.Errorf("Expected UserView id but got %v", err)
		return
	}

	if newUserView.Id != userViewId {
		t.Errorf("Expected NewUserView id to equal UserView id but got %s", newUserView.Id)
	}
}
