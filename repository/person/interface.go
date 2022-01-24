package person

import "Project/research/sample-gql/entities"

type Person interface {
	Get() ([]entities.Person, error)
	Create(entities.Person) (entities.Person, error)
}
