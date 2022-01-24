package book

import (
	"Project/research/sample-gql/entities"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Get() ([]entities.Book, error) {
	var books []entities.Book
	var tmp []entities.Person
	qry := br.db.Raw("Select * from books join people on people.ID = books.author").Scan(&books)
	// br.db.Joins("Persons").Find(&books)

	if err := qry.Error; err != nil {
		return nil, err
	}

	fmt.Println(qry.Statement)
	fmt.Println(qry.Statement.Preloads["Books"]...)
	fmt.Println(tmp)
	return books, nil
}

func (br *BookRepository) Create(book entities.Book) (entities.Book, error) {
	if err := br.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
