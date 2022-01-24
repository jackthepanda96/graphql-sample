package person

import (
	"Project/research/sample-gql/entities"
	"Project/research/sample-gql/repository/person"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PersonController struct {
	repo person.Person
}

func New(repository person.Person) *PersonController {
	return &PersonController{repo: repository}
}

func (pc PersonController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := pc.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Error on server",
				"data":    nil,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    res,
		})
	}
}

func (pc PersonController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := RegisterRequestFormat{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "wrong input format",
				"data":    req,
			})
		}

		res, err := pc.repo.Create(entities.Person{Nama: req.Nama, HP: req.HP, Umur: req.Umur, Password: req.Password})

		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]interface{}{
				"message": "theres is some problem",
				"data":    req,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success register",
			"data":    res,
		})
	}
}
