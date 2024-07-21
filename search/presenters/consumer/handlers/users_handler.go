package handlers

import (
	"encoding/json"
	"fmt"
	"main/search/application/payloads"
	"main/search/application/services"
	"main/search/presenters/consumer/dtos"
)

type UsersHandler[T any] struct {
	SearchService *services.SearchService[T]
}

func (uh *UsersHandler[T]) Create(body []byte) error {
	// Gets body
	var createUserDto dtos.CreateUserDto
	err := json.Unmarshal(body, &createUserDto)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	// Creates user document
	err = uh.SearchService.Create(&payloads.CreatePayload{
		Index:      "users",
		DocumentId: createUserDto.Id,
		Body:       body,
	})

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}

func (uh *UsersHandler[T]) Update(body []byte) error {
	// Gets body
	var updateUserDto dtos.UpdateUserDto
	err := json.Unmarshal(body, &updateUserDto)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	// Updates user document
	err = uh.SearchService.Update(&payloads.UpdatePayload{
		Index:      "users",
		DocumentId: updateUserDto.Id,
		Body:       body,
	})
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}

func (uh *UsersHandler[T]) Delete(body []byte) error {
	// Gets body
	var deleteUserDto dtos.DeleteUserDto
	err := json.Unmarshal(body, &deleteUserDto)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	// Deletes user document
	err = uh.SearchService.Delete(&payloads.DeletePayload{
		Index:      "users",
		DocumentId: deleteUserDto.Id,
	})
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}
