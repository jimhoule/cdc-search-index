package controllers

import (
	"encoding/json"
	"fmt"
	"main/queue"
	"main/router"
	"main/router/utils"
	"main/users/application/payloads"
	"main/users/application/services"
	"main/users/presenters/http/dtos"
	"net/http"
)

type UsersController struct {
	UsersService *services.UsersService
	QueueProducerHandler *queue.ProducerHandler
}

func (uc *UsersController) GetAll(writer http.ResponseWriter, request *http.Request) {
	users, err := uc.UsersService.GetAll()
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, users)
}

func (uc *UsersController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	user, err := uc.UsersService.GetById(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, user)
}

func (uc *UsersController) Create(writer http.ResponseWriter, request *http.Request) {
	var createUserDto dtos.CreateUserDto
	err := utils.ReadHttpRequestBody(writer, request, &createUserDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	user, err := uc.UsersService.Create(&payloads.CreateUserPayload{
		Firstname: createUserDto.Firstname,
		LastName: createUserDto.Lastname,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusCreated, user)

	jsonEncodedUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error sending created user to queue: ", err)
	}

	uc.QueueProducerHandler.SendMessage("users.created", "partition key", jsonEncodedUser)
}

func (uc *UsersController) Update(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	var updateUserDto dtos.UpdateUserDto
	err := utils.ReadHttpRequestBody(writer, request, &updateUserDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	user, err := uc.UsersService.Update(id, &payloads.UpdateUserPayload{
		Firstname: updateUserDto.Firstname,
		LastName: updateUserDto.Lastname,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, user)

	jsonEncodedUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error sending updated user to queue: ", err)
	}

	uc.QueueProducerHandler.SendMessage("users.updated", "partition key", jsonEncodedUser)
}

func (uc *UsersController) Delete(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	_, err := uc.UsersService.Delete(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusNoContent, nil)

	jsonEncodedUserId, err := json.Marshal(id)
	if err != nil {
		fmt.Println("error sending deleted user id to queue: ", err)
	}

	uc.QueueProducerHandler.SendMessage("users.deleted", "partition key", jsonEncodedUserId)
}