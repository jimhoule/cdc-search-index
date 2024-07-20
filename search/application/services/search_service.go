package services

import (
	"main/search/application/payloads"
	"main/search/application/ports"
)

type SearchService struct {
	SearchRepository ports.SearchRepositoryPort
}

func (ss *SearchService) Create(createPayload *payloads.CreatePayload) error {
	return ss.SearchRepository.Create(
		createPayload.Index,
		createPayload.DocumentId,
		createPayload.Body,
	)
}

func (ss *SearchService) Update(updatePayload *payloads.UpdatePayload) error {
	return ss.SearchRepository.Update(
		updatePayload.Index,
		updatePayload.DocumentId,
		updatePayload.Body,
	)
}

func (ss *SearchService) Delete(deletePayload *payloads.DeletePayload) error {
	return ss.SearchRepository.Delete(
		deletePayload.Index,
		deletePayload.DocumentId,
	)
}