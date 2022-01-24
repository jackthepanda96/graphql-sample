package auth

import (
	"Project/research/sample-gql/delivery/common"
	"Project/research/sample-gql/repository/auth"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type AuthController struct {
	repository auth.Auth
}

func New(auth auth.Auth) *AuthController {
	return &AuthController{
		repository: auth,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest LoginRequestFormat

		if err := c.Bind(&loginRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		token, err := a.repository.Login(loginRequest.Username, loginRequest.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		loginResponse := LoginResponseFormat{
			Token: token,
		}

		return c.JSON(http.StatusOK, loginResponse)
	}
}
