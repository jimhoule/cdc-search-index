package services

import (
	"encoding/json"
	"main/search/application/payloads"
	"main/search/domain/views"
	"main/search/infrastructures/persistence/repositories"
	"testing"
)

func getTestContext() (*SearchService[views.UserView], *views.UserView, func(), func() (error)) {
	searchService := &SearchService[views.UserView]{
		SearchRepository: &repositories.FakeSearchRepository[views.UserView]{},
	}

	userView := &views.UserView{
		Id: "dummyDocumentId",
		Firstname: "dummy firstname",
		Lastname: "dummy lastname",
	}
	body, _ := json.Marshal(userView)

	create := func() (error) {
		return searchService.Create(&payloads.CreatePayload{
			Index: "users",
			DocumentId: userView.Id,
			Body: body,
		})
	}

	return searchService, userView, repositories.ResetFakeSearchRepository, create
}

func TestCreateSearchService(t *testing.T) {
	_, _, reset, create := getTestContext()
	defer reset()

	err := create()
	if err != nil {
		t.Errorf("Expected to create a Document but got %v", err)
		return
	}
}

func TestGetDocumentByIdService(t *testing.T) {
	searchService, userView, reset, create := getTestContext()
	defer reset()

	create()

	view, err := searchService.GetByDocumentId(&payloads.GetByDocumentIdPayload{
		Index: "users",
		DocumentId: userView.Id,
	})
	if err != nil {
		t.Errorf("Expected to get a User View by document id but got %v", err)
		return
	}

	if *userView != *view {
		t.Errorf("Expected View to be equal to User View but got %v", view)
		return
	}
}

func TestUpdateSearchService(t *testing.T) {
	searchService, userView, reset, create := getTestContext()
	defer reset()

	create()

	updatedUserView := &views.UserView{
		Id: userView.Id,
		Firstname: "Updated dummy firstname",
		Lastname: userView.Lastname,
	}
	body, _ := json.Marshal(updatedUserView)

	err := searchService.Update(&payloads.UpdatePayload{
		Index: "users",
		DocumentId: userView.Id,
		Body: body,
	})
	if err != nil {
		t.Errorf("Expected User but got %v", err)
		return
	}
}

func TestDeleteUserService(t *testing.T) {
	searchService, userView, reset, create := getTestContext()
	defer reset()

	create()

	err := searchService.Delete(&payloads.DeletePayload{
		Index: "users",
		DocumentId: userView.Id,
	})
	if err != nil {
		t.Errorf("Expected User id but got %v", err)
		return
	}
}