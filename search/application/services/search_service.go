package services

import (
	"main/search/application/payloads"
	"main/search/application/ports"
)

type SearchService[T any] struct {
	SearchRepository ports.SearchRepositoryPort[T]
}

func (ss *SearchService[T]) GetByDocumentId(getByDocumentIdPayload *payloads.GetByDocumentIdPayload) (*T, error) {
	body, err := ss.SearchRepository.GetByDocumentId(
		getByDocumentIdPayload.Index,
		getByDocumentIdPayload.DocumentId,
	)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (ss *SearchService[T]) Create(createPayload *payloads.CreatePayload) (*T, error) {
	return ss.SearchRepository.Create(
		createPayload.Index,
		createPayload.DocumentId,
		createPayload.Body,
	)
}

func (ss *SearchService[T]) Update(updatePayload *payloads.UpdatePayload) (*T, error) {
	return ss.SearchRepository.Update(
		updatePayload.Index,
		updatePayload.DocumentId,
		updatePayload.Body,
	)
}

func (ss *SearchService[T]) Delete(deletePayload *payloads.DeletePayload) (string, error) {
	return ss.SearchRepository.Delete(
		deletePayload.Index,
		deletePayload.DocumentId,
	)
}
