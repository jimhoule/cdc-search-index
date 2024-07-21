package controllers

import (
	"main/router"
	"main/router/utils"
	"main/search/application/payloads"
	"main/search/application/services"
	"net/http"
)

type SearchController[T any] struct {
	SearchService *services.SearchService[T]
}

func (usc *SearchController[T]) GetByDocumentId(writer http.ResponseWriter, request *http.Request) {
	index := router.GetUrlParam(request, "index")
	documentId := router.GetUrlParam(request, "documentId")

	body, err := usc.SearchService.GetByDocumentId(&payloads.GetByDocumentIdPayload{
		Index:      index,
		DocumentId: documentId,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, body)
}
