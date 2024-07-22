package factories

import (
	"main/users/domain/models"
	"main/uuid/services"
)

type UsersFactory struct {
	UuidService services.UuidService
}

func (uf *UsersFactory) Create(firstname string, lastname string) *models.User {
	return &models.User{
		Id:        uf.UuidService.Generate(),
		Firstname: firstname,
		Lastname:  lastname,
	}
}
