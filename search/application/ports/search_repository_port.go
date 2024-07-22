package ports

type SearchRepositoryPort[T any] interface {
	GetByDocumentId(index string, documentId string) (*T, error)
	Create(index string, id string, body []byte) (*T, error)
	Update(index string, documentId string, body []byte) (*T, error)
	Delete(index string, documentId string) (string, error)
}
