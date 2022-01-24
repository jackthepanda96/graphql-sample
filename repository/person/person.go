package person

import (
	"Project/research/sample-gql/entities"

	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (p *PersonRepository) Get() ([]entities.Person, error) {
	var persons []entities.Person
	if err := p.db.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (p *PersonRepository) Create(person entities.Person) (entities.Person, error) {
	if err := p.db.Save(&person).Error; err != nil {
		return person, err
	}
	return person, nil
}
