package services

import (
	"main/users/application/payloads"
	"main/users/application/ports"
	"main/users/domain/factories"
	"main/users/domain/models"
)

type UsersService struct {
	UsersFactory *factories.UsersFactory
	UsersRepository ports.UsersRepositoryPorts
}

func (us *UsersService) GetAll() ([]*models.User, error) {
	return us.UsersRepository.GetAll()
}

func (us *UsersService) GetById(id string) (*models.User, error) {
	return us.UsersRepository.GetById(id)
}

func (us *UsersService) Create(createUserPayload *payloads.CreateUserPayload) (*models.User, error) {
	user := us.UsersFactory.Create(
		createUserPayload.Firstname,
		createUserPayload.LastName,
	)

	return us.UsersRepository.Create(user)
}

func (us *UsersService) Update(id string, updateUserPayload *payloads.UpdateUserPayload) (*models.User, error) {
	updatedUser := &models.User{
		Id: id,
		Firstname: updateUserPayload.Firstname,
		LastName: updateUserPayload.LastName,
	}

	return us.UsersRepository.Update(updatedUser)
}

func (us *UsersService) Delete(id string) (string, error) {
	return us.UsersRepository.Delete(id)
}