package book

import "Project/research/sample-gql/entities"

type Book interface {
	Get() ([]entities.Book, error)
	Create(entities.Book) (entities.Book, error)
}
