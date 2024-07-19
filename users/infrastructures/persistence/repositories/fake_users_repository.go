package repositories

import (
	"fmt"
	"main/users/domain/models"
)

type FakeUsersRepository struct{}

var users []*models.User = []*models.User{}

func ResetFakeUsersRepository() {
	users = []*models.User{}
}

func (*FakeUsersRepository) GetAll() ([]*models.User, error) {
	return users, nil
}

func (*FakeUsersRepository) GetById(id string) (*models.User, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, nil
}

func (*FakeUsersRepository) Create(user *models.User) (*models.User, error) {
	users = append(users, user)

	return user, nil
}

func (*FakeUsersRepository) Update(updatedUser *models.User) (*models.User, error) {
	for index, user := range users {
		if user.Id == updatedUser.Id {
			users[index] = updatedUser

			return updatedUser, nil
		}
	}

	return nil, fmt.Errorf("user with id %s does not exist", updatedUser.Id)
}

func (*FakeUsersRepository) Delete(id string) (string, error) {
	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index + 1:]...)

			return id, nil
		}
	}

	return "", fmt.Errorf("user with id %s does not exist", id)
}