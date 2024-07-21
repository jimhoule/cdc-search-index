package ports

type SearchRepositoryPort[T any] interface {
	GetByDocumentId(index string, documentId string) (*T, error)
	Create(index string, id string, body []byte) error
	Update(index string, documentId string, body []byte) error
	Delete(index string, documentId string) error
}
