package book

import (
	"Project/research/sample-gql/entities"
	"Project/research/sample-gql/entities/model"
)

type Book interface {
	Get() ([]entities.Book, error)
	Create(entities.Book) (entities.Book, error)
	GraphGet() ([]*model.Book, error)
}
