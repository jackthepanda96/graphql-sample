package book

import (
	"Project/research/sample-gql/delivery/common"
	"Project/research/sample-gql/entities"
	"Project/research/sample-gql/repository/book"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	repository book.Book
}

func New(book book.Book) *BookController {
	return &BookController{
		repository: book,
	}
}

func (bc BookController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bc.repository.Get()

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get books",
			"data":    res,
		})
	}
}

func (bc BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bookRequest BookRequestFormat

		if err := c.Bind(&bookRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		response, err := bc.repository.Create(entities.Book{Title: bookRequest.Title, Author: uint(bookRequest.Author)})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, response)
	}
}
