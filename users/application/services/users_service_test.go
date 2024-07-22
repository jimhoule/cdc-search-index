package services

import (
	"main/users/application/payloads"
	"main/users/domain/factories"
	"main/users/domain/models"
	"main/users/infrastructures/persistence/repositories"
	"main/uuid"
	"testing"
)

func getTestContext() (*UsersService, func(), func() (*models.User, error)) {
	usersService := &UsersService{
		UsersFactory: &factories.UsersFactory{
			UuidService: uuid.GetService(),
		},
		UsersRepository: &repositories.FakeUsersRepository{},
	}

	createUser := func() (*models.User, error) {
		return usersService.Create(&payloads.CreateUserPayload{
			Firstname: "Dummy firstname",
			Lastname:  "Dummy lastname",
		})
	}

	return usersService, repositories.ResetFakeUsersRepository, createUser
}

func TestCreateUserService(t *testing.T) {
	_, reset, createUser := getTestContext()
	defer reset()

	_, err := createUser()
	if err != nil {
		t.Errorf("Expected to create a User but got %v", err)
		return
	}
}

func TestGetAllUsersService(t *testing.T) {
	usersService, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	users, err := usersService.GetAll()
	if err != nil {
		t.Errorf("Expected to find all Users but got %v", err)
		return
	}

	if len(users) != 1 {
		t.Errorf("Expected slice of Users with a length of 1 but got %d", len(users))
		return
	}

	if users[0] != newUser {
		t.Errorf("Expected first User of slice to be equal to New User but got %v", users[0])
		return
	}
}

func TestGetUserByIdService(t *testing.T) {
	usersService, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	user, err := usersService.GetById(newUser.Id)
	if err != nil {
		t.Errorf("Expected to get a User by id but got %v", err)
		return
	}

	if newUser != user {
		t.Errorf("Expected User to be equal to New User but got %v", user)
		return
	}
}

func TestUpdateUserService(t *testing.T) {
	usersService, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	updatedFirstname := "Updated fake firstname"
	user, err := usersService.Update(newUser.Id, &payloads.UpdateUserPayload{
		Firstname: updatedFirstname,
		Lastname:  newUser.Lastname,
	})
	if err != nil {
		t.Errorf("Expected User but got %v", err)
		return
	}

	if newUser.Firstname != updatedFirstname {
		t.Errorf("Expected New User firstname to equal updated firstname but got %s", newUser.Firstname)
		return
	}

	if user.Firstname != updatedFirstname {
		t.Errorf("Expected User firstname to equal updated firstname but got %s", user.Firstname)
	}
}

func TestDeleteUserService(t *testing.T) {
	usersService, reset, createUser := getTestContext()
	defer reset()

	newUser, _ := createUser()

	userId, err := usersService.Delete(newUser.Id)
	if err != nil {
		t.Errorf("Expected User id but got %v", err)
		return
	}

	if newUser.Id != userId {
		t.Errorf("Expected New User id to equal User id but got %s", newUser.Id)
	}
}
