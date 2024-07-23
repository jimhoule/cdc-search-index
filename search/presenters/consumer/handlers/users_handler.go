package handlers

import (
	"encoding/json"
	"fmt"
	"main/search/application/payloads"
	"main/search/application/services"
	"main/search/domain/indices"
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
	_, err = uh.SearchService.Create(&payloads.CreatePayload{
		Index:      indices.UsersIndex,
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
	_, err = uh.SearchService.Update(&payloads.UpdatePayload{
		Index:      indices.UsersIndex,
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
	var deletedUserId string
	err := json.Unmarshal(body, &deletedUserId)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	// Deletes user document
	_, err = uh.SearchService.Delete(&payloads.DeletePayload{
		Index:      indices.UsersIndex,
		DocumentId: deletedUserId,
	})
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}
