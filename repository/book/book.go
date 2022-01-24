package book

import (
	"Project/research/sample-gql/entities"
	"Project/research/sample-gql/entities/model"

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

	qry := br.db.Find(&books)
	// br.db.Raw("Select * from books join people on people.ID = books.author").Scan(&books)
	// br.db.Joins("Persons").Find(&books)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br *BookRepository) Create(book entities.Book) (entities.Book, error) {
	if err := br.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (br *BookRepository) GraphGet() ([]*model.Book, error) {
	type res struct {
		BookID   int
		Title    string
		PersonID string
		Nama     string
		Hp       string
		Umur     int
		Password string
	}
	var tmp []res
	qry := br.db.Raw("Select * from books join people on people.ID = books.author").Scan(&tmp)
	// br.db.Find(&tmp)
	if err := qry.Error; err != nil {
		return nil, err
	}
	resArr := []*model.Book{}
	for i := 0; i < len(tmp); i++ {
		resArr = append(resArr, &model.Book{ID: &tmp[i].BookID,
			Title: tmp[i].Title,
			Author: &model.Person{
				ID:   tmp[i].PersonID,
				Nama: tmp[i].Nama,
				Hp:   &tmp[i].Hp,
				Umur: tmp[i].Umur}})
	}

	return resArr, nil
}
