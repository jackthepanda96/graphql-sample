package router

import (
	"Project/research/sample-gql/delivery/controllers/auth"
	"Project/research/sample-gql/delivery/controllers/book"
	"Project/research/sample-gql/delivery/controllers/person"
	custom "Project/research/sample-gql/delivery/middleware"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id int
}

type ContextKey struct {
	Name string
}

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
			fmt.Println(c.Get("ID"))
			// tmpUser := &User{c.Get("ID").(int)}
			// newReq := c.Request().WithContext(context.WithValue(c.Request().Context(), &ContextKey{"user"}, &tmpUser))
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		}, custom.JWTMiddleware())

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
