package ports

import "main/users/domain/models"

type UsersRepositoryPorts interface {
	GetAll() ([]*models.User, error)
	GetById(id string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(updatedUser *models.User) (*models.User, error)
	Delete(id string) (string, error)
}