package services

import (
	"main/search/application/payloads"
	"main/search/application/ports"
)

type SearchService[T any] struct {
	SearchRepository ports.SearchRepositoryPort[T]
}

func (ss *SearchService[T]) GetAllByIndex(getAllByIndexPayload *payloads.GetByAllByIndexPayload) ([]*T, error) {
	return ss.SearchRepository.GetAllByIndex(getAllByIndexPayload.Index)
}

func (ss *SearchService[T]) GetByDocumentId(getByDocumentIdPayload *payloads.GetByDocumentIdPayload) (*T, error) {
	return ss.SearchRepository.GetByDocumentId(
		getByDocumentIdPayload.Index,
		getByDocumentIdPayload.DocumentId,
	)
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
