package ports

type SearchRepositoryPort interface {
	Create(index string, id string, body []byte) error
	Update(index string, documentId string, body []byte) error
	Delete(index string, documentId string) error
}