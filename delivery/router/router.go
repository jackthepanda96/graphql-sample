package router

import (
	"Project/research/sample-gql/delivery/controllers/auth"
	"Project/research/sample-gql/delivery/controllers/book"
	"Project/research/sample-gql/delivery/controllers/person"
	custom "Project/research/sample-gql/delivery/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, authController *auth.AuthController, personController *person.PersonController, bookController *book.BookController, srv *handler.Server) {
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// Login
	e.POST("/login", authController.Login())
	e.POST("/persons", personController.Create(), custom.JWTMiddleware())
	e.GET("/persons", personController.Get(), custom.JWTMiddleware())
	e.POST("/books", bookController.Create())
	e.GET("/books", bookController.Get())
	{
		e.Use(middleware.CORSWithConfig((middleware.CORSConfig{})))

		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		// For Subscriptions
		e.GET("/subscriptions", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

}
